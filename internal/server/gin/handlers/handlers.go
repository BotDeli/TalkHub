package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
)

func SetHandlers(router *gin.Engine, host string, displayA accountControl.Display, displayU userController.Display, displayM meetingController.Display) {
	router.GET("/", handlerShowMainPage(displayA))
	router.GET("/registration", handlerShowRegistrationPage(displayA))
	router.GET("/hub", handlerShowHubPage(displayU))
	router.GET("/settings", handlerShowSettingsPage(displayU))

	router.GET("/logOut", handlerLogOut(host, displayA))
	router.POST("/createAccount", handlerSignUp(host, displayA, displayU))
	router.POST("/goToAccount", handlerSignIn(host, displayA))

	router.GET("/setEnLanguage", handlerSetEnLanguage(host))
	router.GET("/setRuLanguage", handlerSetRuLanguage(host))

	router.POST("/createMeeting", handlerCreateMeeting(displayM))
	router.GET("/getMyMeetings", handlerGetMyMeetings(displayM))
}
