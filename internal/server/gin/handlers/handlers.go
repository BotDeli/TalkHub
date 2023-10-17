package handlers

import (
	"TalkHub/internal/api/authorization"
	"github.com/gin-gonic/gin"
)

func SetHandlers(router *gin.Engine, display authorization.Display, host string) {
	router.GET("/", handlerShowMainPage(display))
	router.GET("/registration", handlerShowRegistrationPage(display))

	router.POST("goToAccount", handlerSignIn(display, host))
	router.POST("createAccount", handlerSignUp(display, host))

	router.DELETE("/exitAccount", handlerExitAccount(host))
}
