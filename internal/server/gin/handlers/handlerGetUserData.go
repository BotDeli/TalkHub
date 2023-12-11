package handlers

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerGetUserData(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := "Guest"

		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			id = guestID
			guestID++
		} else {
			user, err := displayU.GetUserInfoFromID(id)
			if err == nil {
				username = user.FirstName + " " + user.LastName
			} else {
				guestID++
				id = guestID
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"userID":   id,
		})
	}
}

func handlerGetFullUserData(displayU userController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		user, err := displayU.GetUserInfoFromID(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
