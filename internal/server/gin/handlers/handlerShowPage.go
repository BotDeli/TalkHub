package handlers

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerShowMainPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if value, _ := ctx.Get("lang"); value == "ru" {
			redirectAuthorizedUsers(ctx, "ru-main.html", nil)
		} else {
			redirectAuthorizedUsers(ctx, "en-main.html", nil)
		}
	}
}

func redirectAuthorizedUsers(ctx *gin.Context, nameHtml string, obj any) {
	if value, _ := ctx.Get("id"); value == "" {
		ctx.HTML(http.StatusOK, nameHtml, obj)
	} else {
		ctx.Redirect(http.StatusPermanentRedirect, "/hub")
	}
}

func handlerShowRegistrationPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if value, _ := ctx.Get("lang"); value == "ru" {
			redirectAuthorizedUsers(ctx, "ru-registration.html", nil)
		} else {
			redirectAuthorizedUsers(ctx, "en-registration.html", nil)
		}
	}
}

func handlerShowHubPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		user, err := displayU.GetUserInfoFromID(id)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		ctx.HTML(200, "hub.html", gin.H{
			"Username": user.FirstName + " " + user.LastName,
		})
	}
}

func handlerShowSettingsPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := getUserID(ctx)
		//if id == "" {
		//	ctx.Status(http.StatusNotFound)
		//	return
		//}
		//
		//user, err := displayU.GetUserInfoFromID(id)
		//if err != nil {
		//	ctx.Status(http.StatusNotFound)
		//	return
		//}
		//
		//ctx.HTML(200, "profile.html", gin.H{
		//	"Username": user.FirstName + " " + user.LastName,
		//	"UserID":   id,
		//})
	}
}

var guestID = 1

func handlerShowMeetingPage(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingID := params.GetParamsMeetingId(ctx, displayM)
		if meetingID == "" {
			return
		}

		ctx.HTML(http.StatusOK, "meeting.html", gin.H{
			"NumberRoom": meetingID,
		})
	}
}
