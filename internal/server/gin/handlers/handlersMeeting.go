package handlers

import (
	"TalkHub/internal/server/gin/context"
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/pkg/decoder"
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
		id := context.GetUserIDFromContext(ctx)
		if id == nil {
			return
		}

		meeting, err := decoder.JSONDecoder[MeetingID](ctx.Request.Body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
		}

		displayM.StartMeeting(id.(string), meeting.Id)
		ctx.Status(http.StatusAccepted)
	}
}
