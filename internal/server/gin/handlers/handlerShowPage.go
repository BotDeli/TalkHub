package handlers

import (
	"TalkHub/internal/api/authorization"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerShowMainPage(display authorization.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !showAuthorizedPage(ctx, display) {
			ctx.HTML(http.StatusOK, "main.html", nil)
		}
	}
}

func showAuthorizedPage(ctx *gin.Context, display authorization.Display) (authorized bool) {
	key, err := ctx.Cookie("session")
	if err == nil {
		if /*login*/ _, strErr := display.IsAuthenticated(key); strErr == "" {
			ctx.HTML(http.StatusOK, "mainAuthorized.html", nil)
			return true
		}
	}
	return false
}

func handlerShowRegistrationPage(display authorization.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !showAuthorizedPage(ctx, display) {
			ctx.HTML(http.StatusOK, "registration.html", nil)
		}
	}
}
