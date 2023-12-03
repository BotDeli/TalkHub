package sockets

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/handlers/sockets/hub"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func handlerStreamSocket(displayU userController.Display, displayM meetingController.Display, displayH hub.Display) gin.HandlerFunc {
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
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
				return
			}

			username = user.FirstName + " " + user.LastName
		}

		var client *hub.Client

		if userID == nil {
			client = getGuestClient(meetingID, conn)
		} else {
			client = getAuthorizedClient(userID, username, meetingID, conn)
		}

		time.Sleep(time.Second)

		err := displayH.ConnectToMeeting(client)
		if err != nil {
			ctx.Status(http.StatusLocked)
			return
		}

		var msg hub.StreamMessage

		for {
			err = conn.ReadJSON(&msg)
			if err != nil {
				displayH.DisconnectFromMeeting(client)
				break
			}

			if msg.Recipient == "" {
				displayH.InformAllClients(client, msg)
			} else {
				displayH.InformSpecificClient(client, msg)
			}
		}
	}
}

func getGuestClient(meetingID string, conn *websocket.Conn) *hub.Client {
	return &hub.Client{
		UserID:    "",
		Username:  "",
		MeetingID: meetingID,
		Guest:     true,
		Conn:      conn,
	}
}

func getAuthorizedClient(UserID any, Username, meetingID string, conn *websocket.Conn) *hub.Client {
	return &hub.Client{
		UserID:    UserID,
		Username:  Username,
		MeetingID: meetingID,
		Guest:     false,
		Conn:      conn,
	}
}
