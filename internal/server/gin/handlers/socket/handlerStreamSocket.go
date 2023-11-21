package socket

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/pkg/generator"
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

func handlerStreamSocket(displayU userController.Display, displayM meetingController.Display) gin.HandlerFunc {
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

		var username string
		userID := context.GetUserIDFromContext(ctx)

		if userID != nil {
			user, err := displayU.GetUserInfoFromID(userID)

			if err == nil {
				username = user.FirstName + " " + user.LastName
			}

		} else {
			userID = getTempUserID()
		}

		if username == "" {
			username = "Guest"
		}

		err := NotifyConnect(conn, meetingID, username, userID, displayM)

		if err != nil {
			ctx.Status(http.StatusLocked)
			return
		}

		go AwaitNotifyDisconnect(conn, meetingID, username, userID, displayM)

		var msg StreamMessage

		for {
			err = conn.ReadJSON(&msg)
			if err != nil {
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

func getTempUserID() string {
	return generator.NewUUIDDigitsLetters()
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

func AwaitNotifyDisconnect(conn *websocket.Conn, meetingID, username string, userID any, displayM meetingController.Display) {
	var err error
	for err == nil {
		time.Sleep(5 * time.Second)
		err = conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(5*time.Second))
	}

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
}
