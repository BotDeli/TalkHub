package meetingController

import "time"

type Display interface {
	CreateNewMeeting(OwnerUserID string, name string, date time.Time) error
	GetMyMeetings(userID string) []Meeting
	StartMeeting(ownerUserID, meetingID string) error
	EndMeeting(ownerUserID, meetingID string) error
	ConnectToMeeting(meetingID string) error
	DisconnectToMeeting(meetingID string)
	isActiveMeeting(meetingID string) bool
}

type Meeting struct {
	MeetingID      string
	Name           string
	Date           time.Time
	Started        bool
	CountConnected int
}
