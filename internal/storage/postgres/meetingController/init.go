package meetingController

import "time"

type Display interface {
	CreateNewMeeting(OwnerUserID string, name string, date time.Time) error
	GetMyMeetings(ownerUserID string) []Meeting
	StartMeeting(ownerUserID, meetingID string)
	EndMeeting(ownerUserID, meetingID string)
	ConnectToMeeting(meetingID string) error
	DisconnectToMeeting(meetingID string)
	IsStartedMeeting(meetingID string) bool
}

type Meeting struct {
	MeetingID      string
	Name           string
	Date           time.Time
	Started        bool
	CountConnected int
}
