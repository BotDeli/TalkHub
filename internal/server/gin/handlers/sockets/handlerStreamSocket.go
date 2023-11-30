package sockets

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/internal/tempStorage/tempUserID"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type StreamMessage struct {
	Sender         string `json:"sender"`
	SenderUsername string `json:"username"`
	Recipient      string `json:"recipient"`
	Action         string `json:"action"`
	Data           string `json:"data"`
}

var streamConnections = make(map[string]map[any]*websocket.Conn)

func handlerStreamSocket(displayU userController.Display, displayTU tempUserID.Display, displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingID := params.GetParamsMeetingId(ctx, displayM)
		if meetingID == "" {
			return
		}

		conn := upgradeContextToSocketConnection(ctx)
		if conn == nil {
			return
		}

		defer closeSocketConnection(conn)

		var (
			username string
			err      error
		)
		userID := context.GetUserIDFromContext(ctx)

		if userID != nil {
			user, err := displayU.GetUserInfoFromID(userID)

			if err == nil {
				username = user.FirstName + " " + user.LastName
			}

		} else {
			userID, err = displayTU.TakeTempUserID()
			if err != nil {
				ctx.Status(http.StatusConflict)
				return
			}
		}

		if username == "" {
			username = "Guest"
		}

		time.Sleep(time.Second)

		err = NotifyConnect(conn, meetingID, username, userID, displayM)

		if err != nil {
			ctx.Status(http.StatusLocked)
			return
		}

		var msg StreamMessage

		for {
			err = conn.ReadJSON(&msg)
			if err != nil {
				NotifyDisconnect(meetingID, username, userID, displayTU, displayM)
				break
			}

			if msg.Recipient != "" {
				if otherConn, ok := streamConnections[meetingID][msg.Recipient]; ok {
					_ = otherConn.WriteJSON(msg)
				}
			} else {
				for otherUserID, otherConn := range streamConnections[meetingID] {
					if otherUserID != userID {
						_ = otherConn.WriteJSON(msg)
					}
				}
			}
		}
	}
}

func NotifyConnect(conn *websocket.Conn, meetingID, username string, userID any, displayM meetingController.Display) error {
	err := displayM.ConnectToMeeting(meetingID)
	if err != nil {
		return err
	}

	addConnToStreamConnections(conn, meetingID, userID)

	for otherUserID, otherConn := range streamConnections[meetingID] {
		if otherUserID != userID {
			_ = otherConn.WriteJSON(StreamMessage{
				Sender:         userID.(string),
				SenderUsername: username,
				Recipient:      "",
				Action:         "1",
				Data:           "",
			})
		}
	}

	return nil
}

func addConnToStreamConnections(conn *websocket.Conn, meetingID string, userID any) {
	if _, ok := streamConnections[meetingID]; ok {
		streamConnections[meetingID][userID] = conn
	} else {
		streamConnections[meetingID] = map[any]*websocket.Conn{
			userID: conn,
		}
	}
}

func NotifyDisconnect(meetingID, username string, userID any, displayTU tempUserID.Display, displayM meetingController.Display) {
	delete(streamConnections[meetingID], userID)

	for _, otherConn := range streamConnections[meetingID] {
		_ = otherConn.WriteJSON(StreamMessage{
			Sender:         userID.(string),
			SenderUsername: username,
			Recipient:      "",
			Action:         "0",
			Data:           "",
		})
	}

	displayM.DisconnectFromMeeting(meetingID)
	displayTU.GiveTempUserID(userID)
}
