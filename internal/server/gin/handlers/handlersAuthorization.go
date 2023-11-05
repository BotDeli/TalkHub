package handlers

import (
	"TalkHub/internal/api/accountControl"
	"TalkHub/internal/api/accountControl/pb"
	"TalkHub/internal/server/gin/cookie"
	"TalkHub/internal/storage/postgres/userInfo"
	"TalkHub/pkg/decoder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handlerSignUp(host string, displayA accountControl.Display, displayU userInfo.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := getSignUpData(ctx)
		if data == nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		session := accountAuthorization(ctx, host, data.Email, data.Password, displayA.Registration)

		if session == nil {
			ctx.Status(http.StatusBadRequest)
		}

		displayU.SaveUserInfo(&userInfo.User{
			Id:        session.Id,
			UserIcon:  "",
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
		})
		ctx.Status(http.StatusCreated)
	}
}

func getSignUpData(ctx *gin.Context) *SignUpData {
	data, err := decoder.JSONDecoder[SignUpData](ctx.Request.Body)
	if err != nil || data.FirstName == "" || data.LastName == "" || len(data.Email) < 5 || len(data.Password) < 8 {
		return nil
	}
	return data
}

func accountAuthorization(ctx *gin.Context, host, email, password string, ff func(string, string) (*pb.SessionData, string)) *pb.SessionData {
	session, strErr := ff(email, password)
	if strErr != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": strErr})
		return nil
	}

	cookie.SetSessionKey(ctx, host, session.Key)
	return session
}

func handlerSignIn(host string, displayA accountControl.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := getSignInData(ctx)
		if data == nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		session := accountAuthorization(ctx, host, data.Email, data.Password, displayA.Authorization)
		if session == nil {
			ctx.Status(http.StatusBadRequest)
		}
		ctx.JSON(200, nil)
	}
}

func getSignInData(ctx *gin.Context) *SignInData {
	data, err := decoder.JSONDecoder[SignInData](ctx.Request.Body)
	if err != nil || len(data.Email) < 5 || len(data.Password) < 8 {
		return nil
	}
	return data
}

func handlerExitAccount(host string, displayA accountControl.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := cookie.GetSessionKey(ctx)
		displayA.DeleteSession(key)
		cookie.SetSessionKey(ctx, host, "")
	}
}
