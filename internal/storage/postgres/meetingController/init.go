package meetingController

import "time"

type Display interface {
	CreateMeeting(OwnerUserID string, name string, date time.Time) (string, error)
	GetMyMeetings(ownerUserID string) []Meeting
	StartMeeting(ownerUserID, meetingID string)
	EndMeeting(ownerUserID, meetingID string)
	ConnectToMeeting(meetingID string) error
	DisconnectFromMeeting(meetingID string)
	IsStartedMeeting(meetingID string) bool
	UpdateMeetingName(ownerUserID, meetingID, newName string)
	UpdateMeetingDatetime(ownerUserID, meetingID string, newDate time.Time)
	GetMeetingOwnerID(meetingId string) (string, error)
}

type Meeting struct {
	MeetingID      string    `json:"id"`
	Name           string    `json:"name"`
	Datetime       time.Time `json:"date"`
	Started        bool      `json:"started"`
	CountConnected int       `json:"count_connected"`
}
