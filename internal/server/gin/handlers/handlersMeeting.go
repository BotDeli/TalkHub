package handlers

import (
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
		id := getUserID(ctx)
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
		id := getUserID(ctx)
		if id == nil {
			return
		}

		meetings := displayM.GetMyMeetings(id.(string))
		ctx.JSON(http.StatusOK, gin.H{
			"meetings": meetings,
		})
	}
}
