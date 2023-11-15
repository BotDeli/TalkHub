package handlers

import (
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections = make(map[string][]*websocket.Conn)

func handlerChatWebsocket(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingID := params.GetParamsMeetingId(ctx, displayM)
		if meetingID == "" {
			return
		}

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		defer conn.Close()

		if _, ok := connections[meetingID]; ok {
			connections[meetingID] = append(connections[meetingID], conn)
		} else {
			connections[meetingID] = []*websocket.Conn{conn}
		}

		for {
			var msg Message
			err = conn.ReadJSON(&msg)
			if err != nil {
				log.Println(err)
				break
			}

			for _, otherConn := range connections[meetingID] {
				_ = otherConn.WriteJSON(msg)
			}
		}
	}
}
