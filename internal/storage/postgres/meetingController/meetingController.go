package meetingController

import (
	"TalkHub/internal/storage/postgres"
	"database/sql"
	"log"
	"time"
)

type MCDisplay struct {
	PG *postgres.Storage
}

func InitDisplay(pg *postgres.Storage) Display {
	initTable(pg.DB)
	return &MCDisplay{PG: pg}
}

func initTable(db *sql.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS meetings (
    	id VARCHAR NOT NULL PRIMARY KEY UNIQUE,
    	name VARCHAR NOT NULL,
    	date DATE NOT NULL,
    	started BOOLEAN NOT NULL,
    	count_connected INTEGER NOT NULL
	)`); err != nil {
		log.Printf("Error creating meetings table: %s\n", err)
	}
}

func (m *MCDisplay) CreateNewMeeting(OwnerUserID string, name string, date time.Time) error {
	return nil
}
func (m *MCDisplay) GetMyMeetings(userID string) []Meeting {
	return nil
}
func (m *MCDisplay) StartMeeting(ownerUserID, meetingID string) error {
	return nil
}
func (m *MCDisplay) EndMeeting(ownerUserID, meetingID string) error {
	return nil
}
func (m *MCDisplay) ConnectToMeeting(meetingID string) error {
	return nil
}
func (m *MCDisplay) DisconnectToMeeting(meetingID string) {
	return
}
func (m *MCDisplay) isActiveMeeting(meetingID string) bool {
	return false
}
