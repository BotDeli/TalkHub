package context

import (
	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(ctx *gin.Context) any {
	id, have := ctx.Get("id")
	if !have || id == "" {
		return nil
	}

	return id
}
