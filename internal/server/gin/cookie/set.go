package cookie

import (
	"github.com/gin-gonic/gin"
)

func SetSessionKey(ctx *gin.Context, host, key string) {
	ctx.SetCookie("session", key, 24*3600, "/", host, false, true)
}
