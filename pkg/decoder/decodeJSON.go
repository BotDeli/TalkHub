package decoder

import "github.com/gin-gonic/gin"

func DecodeJSON[T any](ctx *gin.Context) (T, error) {
	var data T
	err := ctx.ShouldBindJSON(&data)
	return data, err
}
