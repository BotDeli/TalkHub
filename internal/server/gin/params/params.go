package params

import (
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/pkg/selector"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParamsMeetingId(ctx *gin.Context, displayM meetingController.Display) string {
	meetingID := ctx.Param("id")
	if meetingID == "" || !displayM.IsStartedMeeting(meetingID) {
		errNameMeeting := selector.SelectLanguageFormat(ctx, "error-meeting.html")
		ctx.HTML(http.StatusNotFound, errNameMeeting, nil)

		return ""
	}
	return meetingID
}
