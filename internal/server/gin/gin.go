package gin

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin/handlers"
	"TalkHub/internal/server/gin/middleware"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
)

func StartGinServer(cfg *config.HttpConfig, webrtcCfg *config.WebrtcConfig, displayA accountControl.Display, displayU userController.Display, displayM meetingController.Display) error {
	router := gin.Default()
	loadAllFiles(router)
	router.Use(middleware.CheckerAuthorizedUser(displayA), middleware.CheckerLanguageSelect(cfg.Host))
	handlers.SetHandlers(router, cfg.Host, webrtcCfg, displayA, displayU, displayM)
	return router.Run(cfg.GetAddress())
}

func loadAllFiles(router *gin.Engine) {
	router.LoadHTMLGlob("template/*/*.html")
	router.Static("/static/", "./static")
	router.Static("/images/", "./images")
}
