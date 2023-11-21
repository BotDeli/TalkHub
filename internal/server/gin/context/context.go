package context

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserIDFromContext(ctx *gin.Context) any {
	id, have := ctx.Get("id")
	if !have || id == "" {
		ctx.Status(http.StatusUnauthorized)
		return nil
	}

	return id
}
