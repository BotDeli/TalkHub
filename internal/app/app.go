package app

import (
	"TalkHub/internal/config"
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"log"
)

func Start() {
	cfg := config.MustReadConfig()
	//displayA := accountControl.MustInitDisplay(cfg.Grpc)

	pg, err := postgres.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	userController.InitDisplay(pg)
	meetingController.InitDisplay(pg)

	//displayU := userController.InitDisplay(pg)
	//log.Fatal(gin.StartGinServer(cfg.Http, displayA, displayU))
}
