package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/config"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
)

func SetHandlers(router *gin.Engine, host string, webrtcCfg *config.WebrtcConfig, displayA accountControl.Display, displayU userController.Display, displayM meetingController.Display) {
	// main pages
	router.GET("/", handlerShowMainPage(displayU))
	router.GET("/registration", handlerShowRegistrationPage(displayU))
	router.GET("/hub", handlerShowHubPage(displayU))
	router.GET("/settings", handlerShowSettingsPage(displayU))

	// functions authorization
	router.GET("/logOut", handlerLogOut(host, displayA))
	router.POST("/createAccount", handlerSignUp(host, displayA, displayU))
	router.POST("/goToAccount", handlerSignIn(host, displayA))

	// language settings
	router.GET("/setEnLanguage", handlerSetEnLanguage(host))
	router.GET("/setRuLanguage", handlerSetRuLanguage(host))

	// functions meeting
	router.POST("/createMeeting", handlerCreateMeeting(displayM))
	router.GET("/getMyMeetings", handlerGetMyMeetings(displayM))
	router.Handle("UPDATE", "/startMeeting", handlerStartMeeting(displayM))
	router.Handle("DELETE", "/cancelMeeting", handlerCancelMeeting(displayM))
	router.Handle("UPDATE", "/changeMeetingName", handlerChangeMeetingName(displayM))
	router.Handle("UPDATE", "/changeMeetingDatetime", handlerChangeMeetingDatetime(displayM))

	// help functions stream socket
	router.GET("/meeting/:id", handlerShowMeetingPage(displayM))
	router.GET("/getUserData", handlerGetUserData(displayU))
	router.GET("/webrtcConfig", handlerGetWebrtcConfig(webrtcCfg))

	// functions settings
	router.POST("/changePassword", handlerChangePassword(displayU, displayA))
}
