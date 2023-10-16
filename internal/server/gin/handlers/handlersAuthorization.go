package handlers

import (
	"TalkHub/internal/api/authorization"
	"TalkHub/internal/api/authorization/pb"
	"TalkHub/internal/server/gin/cookie"
	"TalkHub/pkg/decoder"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerSignUp(display authorization.Display, host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startAuthorization(ctx, display.Register, host)
	}
}

func handlerSignIn(display authorization.Display, host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startAuthorization(ctx, display.LogIn, host)
	}
}

func handlerExitAccount(host string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie.SetSessionKey(ctx, host, "")
	}
}

func startAuthorization(ctx *gin.Context, ff func(string, string) (string, string), host string) {
	u := getUserData(ctx)
	if u == nil {
		return
	}

	key, strErr := ff(u.Login, u.Password)
	if strErr != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": strErr})
		return
	}

	fmt.Println(key)

	cookie.SetSessionKey(ctx, host, key)
	ctx.JSON(http.StatusOK, gin.H{"error": ""})
}

func getUserData(ctx *gin.Context) *pb.User {
	u, err := decoder.DecodeJSON[pb.User](ctx)
	if err != nil || u.Login == "" || u.Password == "" {
		ctx.JSON(http.StatusBadRequest, nil)
		return nil
	}
	return &u
}
