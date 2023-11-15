package params

import (
	"TalkHub/internal/storage/postgres/meetingController"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParamsMeetingId(ctx *gin.Context, displayM meetingController.Display) string {
	meetingID := ctx.Param("id")
	if meetingID == "" || !displayM.IsStartedMeeting(meetingID) {
		ctx.HTML(http.StatusNotFound, "error-meeting.html", nil)
		return ""
	}
	return meetingID
}
