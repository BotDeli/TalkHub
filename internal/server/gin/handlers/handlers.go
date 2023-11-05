package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/storage/postgres/userInfo"
	"github.com/gin-gonic/gin"
)

func SetHandlers(router *gin.Engine, host string, displayA accountControl.Display, displayU userInfo.Display) {
	router.GET("/", handlerShowMainPage(displayA))
	router.GET("/registration", handlerShowRegistrationPage(displayA))
	router.GET("/hub", handlerShowHubPage())

	router.POST("/createAccount", handlerSignUp(host, displayA, displayU))
	router.POST("/goToAccount", handlerSignIn(host, displayA))

	router.DELETE("/exitAccount", handlerExitAccount(host, displayA))
}
