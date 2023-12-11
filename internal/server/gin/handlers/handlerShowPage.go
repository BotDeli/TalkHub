package handlers

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/pkg/selector"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerShowMainPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nameHtml := selector.SelectLanguageFormat(ctx, "main.html")
		redirectAuthorizedUsers(ctx, nameHtml, nil, displayU)
	}
}

func redirectAuthorizedUsers(ctx *gin.Context, nameHtml string, obj any, displayU userController.Display) {
	id := context.GetUserIDFromContext(ctx)
	if id != nil {
		user, err := displayU.GetUserInfoFromID(id)
		if err == nil && user != nil {
			ctx.Redirect(http.StatusPermanentRedirect, "/hub")
			return
		}
	}

	ctx.HTML(http.StatusOK, nameHtml, obj)
}

func handlerShowRegistrationPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nameHtml := selector.SelectLanguageFormat(ctx, "registration.html")
		redirectAuthorizedUsers(ctx, nameHtml, nil, displayU)
	}
}

func handlerShowHubPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		user, err := displayU.GetUserInfoFromID(id)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		ctx.HTML(http.StatusOK, "hub.html", gin.H{
			"Username": user.FirstName + " " + user.LastName,
		})
	}
}

func handlerShowSettingsPage(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		user, err := displayU.GetUserInfoFromID(id)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		ctx.HTML(http.StatusOK, "settings.html", gin.H{
			"UserID":    user.Id,
			"FirstName": user.FirstName,
			"LastName":  user.LastName,
			"Email":     user.Email,
		})
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
