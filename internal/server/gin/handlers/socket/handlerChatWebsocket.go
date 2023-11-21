package socket

import (
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgraderChat = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var chatConnections = make(map[string][]*websocket.Conn)

func handlerChatWebsocket(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingID := params.GetParamsMeetingId(ctx, displayM)
		if meetingID == "" {
			return
		}

		conn, err := upgraderChat.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		defer conn.Close()

		if _, ok := chatConnections[meetingID]; ok {
			chatConnections[meetingID] = append(chatConnections[meetingID], conn)
		} else {
			chatConnections[meetingID] = []*websocket.Conn{conn}
		}

		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			for _, otherConn := range chatConnections[meetingID] {
				_ = otherConn.WriteMessage(t, msg)
			}
		}
	}
}
