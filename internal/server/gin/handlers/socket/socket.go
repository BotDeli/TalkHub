package socket

import (
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func SetSocketHandlers(router *gin.Engine, displayU userController.Display, displayM meetingController.Display) {
	router.GET("/meeting/:id/chat", handlerChatWebsocket(displayM))
	router.GET("/meeting/:id/stream", handlerStreamSocket(displayU, displayM))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func upgradeContextToSocketConnection(ctx *gin.Context) *websocket.Conn {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	return conn
}

func closeSocketConnection(conn *websocket.Conn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing socket connection: %v", err)
	}
}
