package handlers

import (
	"TalkHub/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerGetWebrtcConfig(cfg *config.WebrtcConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cfg)
	}
}
