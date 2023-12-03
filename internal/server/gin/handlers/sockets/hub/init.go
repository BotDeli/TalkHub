package hub

import (
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/tempStorage/tempUserID"
	"sync"
)

type Display interface {
	ConnectToMeeting(currentClient *Client) error
	DisconnectFromMeeting(currentClient *Client)
	InformAllClients(currentClient *Client, msg StreamMessage)
	InformSpecificClient(currentClient *Client, msg StreamMessage)
}

func InitDisplay(displayM meetingController.Display, displayTU tempUserID.Display) Display {
	return &Hub{
		Meetings:  make(map[string]Clients),
		displayM:  displayM,
		displayTU: displayTU,
		mutex:     sync.Mutex{},
	}
}
