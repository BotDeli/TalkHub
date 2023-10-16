package gin

import (
	"TalkHub/internal/api/authorization"
	"TalkHub/internal/config"
	"TalkHub/internal/server/gin/handlers"
	"github.com/gin-gonic/gin"
)

func StartGinServer(cfg *config.HttpConfig, display authorization.Display) error {
	router := gin.Default()
	loadAllFiles(router)
	handlers.SetHandlers(router, display, cfg.Host)
	return router.Run(cfg.GetAddress())
}

func loadAllFiles(router *gin.Engine) {
	router.LoadHTMLGlob("template/*")
	router.Static("/static/", "./static")
}
