package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/storage/postgres/userController"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var errNotAuthorized = errors.New("not authorized")

func handlerShowMainPage(displayA accountControl.Display) gin.HandlerFunc {
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

func handlerShowRegistrationPage(displayA accountControl.Display) gin.HandlerFunc {
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
		id := getUserID(ctx)
		if id == nil {
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

func getUserID(ctx *gin.Context) any {
	id, have := ctx.Get("id")
	if !have || id == "" {
		ctx.Status(http.StatusUnauthorized)
		return nil
	}

	return id
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
