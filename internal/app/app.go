package app

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin"
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/internal/tempStorage/tempUserID"
	"log"
)

func Start() {
	cfg := config.MustReadConfig()
	displayA := accountControl.MustInitDisplay(cfg.Grpc)

	pg, err := postgres.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	displayU := userController.InitDisplay(pg)
	displayTU := tempUserID.InitDisplay(displayU)
	displayM := meetingController.InitDisplay(pg, cfg.Meeting)

	log.Fatal(gin.StartGinServer(cfg.Http, cfg.Webrtc, displayA, displayU, displayTU, displayM))
}
