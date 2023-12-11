package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/pkg/decoder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangePasswordForm struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func handlerChangePassword(displayU userController.Display, displayA accountControl.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		form, err := decoder.JSONDecoder[ChangePasswordForm](ctx.Request.Body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		user, err := displayU.GetUserInfoFromID(id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		strErr := displayA.ChangePassword(user.Email, form.Password, form.NewPassword)
		if strErr != "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.Status(http.StatusOK)
	}
}
