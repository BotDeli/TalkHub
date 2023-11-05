package handlers

import (
	"TalkHub/internal/api/accountControl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerShowMainPage(displayA accountControl.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redirectAuthorizedUsers(ctx, "main.html", nil)
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
		redirectAuthorizedUsers(ctx, "registration.html", nil)
	}
}

func handlerShowHubPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redirectDontAuthorizedUsers(ctx, "hub.html", nil)
	}
}

func redirectDontAuthorizedUsers(ctx *gin.Context, nameHtml string, obj any) {
	if value, _ := ctx.Get("id"); value != "" {
		ctx.HTML(http.StatusOK, nameHtml, obj)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
