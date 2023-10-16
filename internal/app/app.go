package app

import (
	"TalkHub/internal/api/authorization"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin"
	"log"
)

func Start() {
	cfg := config.MustReadConfig()
	display := authorization.MustInitDisplay(cfg.Grpc)
	log.Fatal(gin.StartGinServer(cfg.Http, display))
}
