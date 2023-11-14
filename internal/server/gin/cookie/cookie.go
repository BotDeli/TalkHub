package cookie

import "github.com/gin-gonic/gin"

func SetSessionKey(ctx *gin.Context, host, key string) {
	setCookie(ctx, host, "session", key)
}

func setCookie(ctx *gin.Context, host, name, value string) {
	ctx.SetCookie(name, value, 24*3600, "/", host, false, true)
}

func GetSessionKey(ctx *gin.Context) string {
	return getCookie(ctx, "session")
}

func getCookie(ctx *gin.Context, name string) string {
	value, _ := ctx.Cookie(name)
	return value
}

func SetLanguageCookie(ctx *gin.Context, host, lang string) {
	setCookie(ctx, host, "lang", lang)
}

func GetLanguageCookie(ctx *gin.Context) string {
	return getCookie(ctx, "lang")
}
