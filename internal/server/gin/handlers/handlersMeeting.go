package handlers

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/server/gin/params"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/pkg/decoder"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MeetingRequest struct {
	Name string    `json:"name"`
	Date time.Time `json:"datetime"`
}

func handlerCreateMeeting(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		meeting, err := decoder.JSONDecoder[MeetingRequest](ctx.Request.Body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		meetingId, err := displayM.CreateMeeting(id.(string), meeting.Name, meeting.Date)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"id": meetingId,
		})
	}
}

func handlerGetMyMeetings(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		meetings := displayM.GetMyMeetings(id.(string))
		ctx.JSON(http.StatusOK, gin.H{
			"meetings": meetings,
		})
	}
}

type MeetingID struct {
	Id string `json:"id"`
}

func handlerStartMeeting(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingData := getMeetingData(ctx)
		if meetingData == nil {
			return
		}

		displayM.StartMeeting(meetingData.UserID, meetingData.MeetingID)
		ctx.Status(http.StatusAccepted)
	}
}

type MeetingData struct {
	UserID    string
	MeetingID string
}

func getMeetingData(ctx *gin.Context) *MeetingData {
	id := context.GetUserIDFromContext(ctx)
	if id == nil {
		ctx.Status(http.StatusUnauthorized)
		return nil
	}

	meeting, err := decoder.JSONDecoder[MeetingID](ctx.Request.Body)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	return &MeetingData{
		UserID:    id.(string),
		MeetingID: meeting.Id,
	}
}

func handlerCancelMeeting(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		meetingData := getMeetingData(ctx)
		if meetingData == nil {
			return
		}

		displayM.EndMeeting(meetingData.UserID, meetingData.MeetingID)
		ctx.Status(http.StatusAccepted)
	}
}

func handlerEndMeeting(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		meetingID := params.GetParamsMeetingId(ctx, displayM)

		displayM.EndMeeting(id.(string), meetingID)
		ctx.Status(http.StatusAccepted)
	}
}

type RequestChangeMeetingName struct {
	MeetingID string `json:"id"`
	NewName   string `json:"newName"`
}

func handlerChangeMeetingName(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		request, err := decoder.JSONDecoder[RequestChangeMeetingName](ctx.Request.Body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		displayM.UpdateMeetingName(id.(string), request.MeetingID, request.NewName)
		ctx.Status(http.StatusAccepted)
	}
}

type RequestChangeMeetingDatetime struct {
	MeetingID string    `json:"id"`
	NewDate   time.Time `json:"newDatetime"`
}

func handlerChangeMeetingDatetime(displayM meetingController.Display) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		request, err := decoder.JSONDecoder[RequestChangeMeetingDatetime](ctx.Request.Body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		displayM.UpdateMeetingDatetime(id.(string), request.MeetingID, request.NewDate)
		ctx.Status(http.StatusAccepted)
	}
}
