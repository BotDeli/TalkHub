package gin

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin/handlers"
	"TalkHub/internal/server/gin/middleware"
	"TalkHub/internal/storage/postgres/userInfo"
	"github.com/gin-gonic/gin"
)

func StartGinServer(cfg *config.HttpConfig, displayA accountControl.Display, displayU userInfo.Display) error {
	router := gin.Default()
	loadAllFiles(router)
	router.Use(middleware.CheckerAuthorizedUser(displayA))
	handlers.SetHandlers(router, cfg.Host, displayA, displayU)
	return router.Run(cfg.GetAddress())
}

func loadAllFiles(router *gin.Engine) {
	router.LoadHTMLGlob("template/*")
	router.Static("/static/", "./static")
	router.Static("/images/", "./images")
}
