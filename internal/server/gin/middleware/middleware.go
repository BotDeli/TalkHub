package middleware

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/server/gin/cookie"
	"github.com/gin-gonic/gin"
)

func CheckerAuthorizedUser(displayA accountControl.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := cookie.GetSessionKey(ctx)
		if key != "" {
			if id, strErr := displayA.IsAuthenticated(key); strErr == "" {
				ctx.Set("id", id)
				return
			}
		}

		ctx.Set("id", "")
	}
}

func CheckerLanguageSelect(host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := cookie.GetLanguageCookie(ctx)
		if lang == "" {
			cookie.SetLanguageCookie(ctx, host, "en")
			lang = "en"
		}

		ctx.Set("lang", lang)
	}
}
