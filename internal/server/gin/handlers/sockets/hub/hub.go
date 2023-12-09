package hub

import (
	"TalkHub/internal/storage/postgres/meetingController"
	"TalkHub/internal/tempStorage/tempUserID"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type RequestMessage struct {
	Recipient string `json:"recipient"`
	Action    string `json:"action"`
	Data      string `json:"data"`
}

type StreamMessage struct {
	Sender         string `json:"sender"`
	SenderUsername string `json:"username"`
	RequestMessage
}

type Client struct {
	UserID    any
	Username  string
	MeetingID string
	Guest     bool
	Conn      *websocket.Conn
}

type Clients []*Client
type Meetings map[string]Clients

type Hub struct {
	Meetings
	displayM  meetingController.Display
	displayTU tempUserID.Display
	mutex     sync.Mutex
}

func (h *Hub) ConnectToMeeting(currentClient *Client) error {
	err := h.settingGuestClient(currentClient)
	if err != nil {
		return err
	}

	err = h.displayM.ConnectToMeeting(currentClient.MeetingID)
	if err != nil {
		return err
	}

	h.mutex.Lock()
	h.addClientToMeeting(currentClient)
	h.mutex.Unlock()

	h.informAllClientsAboutJoin(currentClient)
	h.Meetings[currentClient.MeetingID] = append(h.Meetings[currentClient.MeetingID], currentClient)

	fmt.Printf("Connected to %s, client %v %s\n", currentClient.MeetingID, currentClient.UserID, currentClient.Username)

	return nil
}

func (h *Hub) settingGuestClient(currentClient *Client) error {
	if currentClient.Guest {
		var err error

		if currentClient.UserID == nil || currentClient.UserID == "" {
			currentClient.UserID, err = h.displayTU.TakeTempUserID()
			if err != nil {
				return err
			}
		}

		if currentClient.Username == "" {
			currentClient.Username = "Guest"
		}
	}
	return nil
}

func (h *Hub) addClientToMeeting(currentClient *Client) {
	clients := h.Meetings[currentClient.MeetingID]
	if clients == nil {
		clients = make([]*Client, 0, 4)
		h.Meetings[currentClient.MeetingID] = clients
	}

	clients = append(clients, currentClient)
}

func (h *Hub) informAllClientsAboutJoin(currentClient *Client) {
	h.InformAllClients(currentClient, StreamMessage{
		Sender:         currentClient.UserID.(string),
		SenderUsername: currentClient.Username,
		RequestMessage: RequestMessage{
			Recipient: "",
			Action:    "1",
			Data:      "",
		},
	})
}

func (h *Hub) InformAllClients(currentClient *Client, msg StreamMessage) {
	clients := h.Meetings[currentClient.MeetingID]
	for _, selectedClient := range clients {
		if selectedClient.UserID != currentClient.UserID {
			_ = selectedClient.Conn.WriteJSON(msg)
		}
	}
}

func (h *Hub) DisconnectFromMeeting(currentClient *Client) {
	h.mutex.Lock()
	h.deleteClientFromMeeting(currentClient)
	h.mutex.Unlock()

	h.informAllClientsAboutLeave(currentClient)

	if currentClient.Guest {
		h.displayTU.GiveTempUserID(currentClient.UserID)
	}

	h.displayM.DisconnectFromMeeting(currentClient.MeetingID)

	fmt.Printf("Disconnected from %s, client %v %s\n", currentClient.MeetingID, currentClient.UserID, currentClient.Username)
}

func (h *Hub) informAllClientsAboutLeave(currentClient *Client) {
	h.InformAllClients(currentClient, StreamMessage{
		Sender:         currentClient.UserID.(string),
		SenderUsername: currentClient.Username,
		RequestMessage: RequestMessage{
			Recipient: "",
			Action:    "0",
			Data:      "",
		},
	})
}

func (h *Hub) deleteClientFromMeeting(currentClient *Client) {
	clients := h.Meetings[currentClient.MeetingID]
	for i, selectedClient := range clients {
		if selectedClient.UserID == currentClient.UserID {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}
}

func (h *Hub) InformSpecificClient(currentClient *Client, msg StreamMessage) {
	clients := h.Meetings[currentClient.MeetingID]
	for _, selectedClient := range clients {
		if selectedClient.UserID == msg.Recipient {
			_ = selectedClient.Conn.WriteJSON(msg)
		}
	}
}

func (h *Hub) ResendRequestMessage(currentClient *Client, rMsg RequestMessage) {
	msg := StreamMessage{
		Sender:         currentClient.UserID.(string),
		SenderUsername: currentClient.Username,
		RequestMessage: rMsg,
	}

	if msg.Recipient == "" {
		h.InformAllClients(currentClient, msg)
	} else {
		h.InformSpecificClient(currentClient, msg)
	}
}
