package app

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin"
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"log"
)

func Start() {
	cfg := config.MustReadConfig()
	displayA := accountControl.MustInitDisplay(cfg.Grpc)

	pg, err := postgres.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	displayU := userController.InitDisplay(pg)
	displayM := meetingController.InitDisplay(pg)

	log.Fatal(gin.StartGinServer(cfg.Http, displayA, displayU, displayM))
}
