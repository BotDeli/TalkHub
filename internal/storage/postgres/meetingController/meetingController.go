package meetingController

import (
	"TalkHub/internal/storage/postgres"
	"TalkHub/pkg/generator"
	"database/sql"
	"errors"
	"log"
	"time"
)

var (
	errMeetingNotCreated  = errors.New("meeting not created")
	errMostCountConnected = errors.New("most connected connections")
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
    	datetime TIMESTAMP NOT NULL,
    	started BOOLEAN NOT NULL,
    	count_connected INTEGER NOT NULL,
    	owner_id VARCHAR NOT NULL
	)`); err != nil {
		log.Printf("Error creating meetings table: %s\n", err)
	}
}

func (m *MCDisplay) CreateMeeting(ownerUserID string, name string, date time.Time) (string, error) {
	query := `INSERT INTO meetings (id, name, datetime, started, count_connected, owner_id) VALUES ($1, $2, $3, $4, $5, $6)`
	id := generator.NewUUIDDigitsLetters()
	_, err := m.PG.DB.Exec(query, id, name, date, false, 0, ownerUserID)
	return id, err
}
func (m *MCDisplay) GetMyMeetings(ownerUserID string) []Meeting {
	query := `SELECT id, name, datetime, started, count_connected FROM meetings WHERE owner_id = $1`
	rows, err := m.PG.DB.Query(query, ownerUserID)
	if err != nil {
		return []Meeting{}
	}

	return scanMeetingsFromRows(rows)
}

func scanMeetingsFromRows(rows *sql.Rows) []Meeting {
	meetings := []Meeting{}
	var (
		meetingID, name string
		datetime        time.Time
		started         bool
		countConnected  int
		err             error
	)

	for rows.Next() {
		err = rows.Scan(&meetingID, &name, &datetime, &started, &countConnected)
		if err != nil {
			continue
		}

		meetings = append(meetings, Meeting{
			MeetingID:      meetingID,
			Name:           name,
			Datetime:       datetime,
			Started:        started,
			CountConnected: countConnected,
		})
	}

	return meetings
}

func (m *MCDisplay) StartMeeting(ownerUserID, meetingID string) {
	query := `UPDATE meetings SET started = 1 WHERE owner_id = $1 AND id = $2`
	_, _ = m.PG.DB.Exec(query, ownerUserID, meetingID)
}

func (m *MCDisplay) EndMeeting(ownerUserID, meetingID string) {
	query := `DELETE FROM meetings WHERE owner_id = $1 AND id = $2`
	_, _ = m.PG.DB.Exec(query, ownerUserID, meetingID)
}

func (m *MCDisplay) ConnectToMeeting(meetingID string) error {
	countConnected, err := getCountConnectedToMeeting(m, meetingID)
	if err != nil {
		return err
	}

	if countConnected >= 12 {
		return errMostCountConnected
	}

	query := `UPDATE meetings SET count_connected = $2 WHERE id = $1`
	_, err = m.PG.DB.Exec(query, meetingID, countConnected+1)
	if err != nil {
		return errMeetingNotCreated
	}
	return nil
}

func getCountConnectedToMeeting(m *MCDisplay, meetingID string) (int, error) {
	query := `SELECT count_connected FROM meetings WHERE id = $1`
	rows, err := m.PG.DB.Query(query, meetingID)
	if err != nil {
		return 0, errMeetingNotCreated
	}

	var countConnected int
	rows.Next()
	err = rows.Scan(&countConnected)
	if err != nil {
		return 0, errMeetingNotCreated
	}

	return countConnected, nil
}

func (m *MCDisplay) DisconnectToMeeting(meetingID string) {
	countConnected, err := getCountConnectedToMeeting(m, meetingID)
	if err != nil {
		return
	}

	if countConnected >= 1 {
		query := `UPDATE meetings SET count_connected = $2 WHERE id = $1`
		_, _ = m.PG.DB.Exec(query, meetingID, countConnected-1)
	}
}

func (m *MCDisplay) IsStartedMeeting(meetingID string) bool {
	query := `SELECT started FROM meetings WHERE id = $1`
	rows, err := m.PG.DB.Query(query, meetingID)
	if err != nil {
		return false
	}

	var started bool
	rows.Next()
	_ = rows.Scan(&started)
	return started
}
