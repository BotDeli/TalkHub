package handlers

import (
	"TalkHub/internal/server/gin/cookie"
	"github.com/gin-gonic/gin"
)

func handlerSetEnLanguage(host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie.SetLanguageCookie(ctx, host, "en")
	}
}

func handlerSetRuLanguage(host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie.SetLanguageCookie(ctx, host, "ru")
	}
}
