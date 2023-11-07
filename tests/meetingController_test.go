package tests

import (
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/meetingController"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

const (
	testUserID    = "10203045fff"
	testMeetingID = "fff1ds020fff"
	testName      = "testMeetingName"
)

var (
	testDate = time.Now()
)

func TestErrorCreateNewMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectExec("INSERT INTO").WithArgs(sqlmock.AnyArg(), testName, testDate, false, 0, testUserID).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.CreateNewMeeting(testUserID, testName, testDate)
		checkErrorIsTestError(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func testingMockMeeting(t *testing.T, initMock func(sqlmock.Sqlmock), testMock func(*testing.T, meetingController.Display)) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	initMock(mockDB)

	display := &meetingController.MCDisplay{PG: &postgres.Storage{DB: db}}

	testMock(t, display)
}

func TestSuccessfulCreateNewMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		result := getEmptyResult()
		mock.ExpectExec("INSERT INTO").WithArgs(sqlmock.AnyArg(), testName, testDate, false, 0, testUserID).WillReturnResult(result).WillReturnError(nil)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.CreateNewMeeting(testUserID, testName, testDate)
		checkErrorIsNil(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestErrorGetMyMeetings(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectQuery("SELECT").WithArgs(testUserID).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		meetings := display.GetMyMeetings(testUserID)
		checkArrayIsEmpty(t, meetings)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestSuccessfulGetMyMeetingsEmptyRows(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := newMeetingRows()
		mock.ExpectQuery("SELECT").WithArgs(testUserID).WillReturnRows(rows)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		meetings := display.GetMyMeetings(testUserID)
		checkArrayIsEmpty(t, meetings)
	}

	testingMockMeeting(t, initMock, testMock)
}

func newMeetingRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "date", "started", "count_connected"})
}

func TestSuccessfulGetMyMeetingsDontEmptyRows(t *testing.T) {
	expectedMeetings := []meetingController.Meeting{
		{testMeetingID, testName, testDate, false, 0},
		{testMeetingID, testName, testDate, true, 4},
		{testMeetingID, testName, testDate, false, 10},
		{testMeetingID, testName, testDate, false, 12},
		{testMeetingID, testName, testDate, true, 2},
		{testMeetingID, testName, testDate, true, 3},
		{testMeetingID, testName, testDate, true, 1},
		{testMeetingID, testName, testDate, false, 9},
	}

	initMock := func(mock sqlmock.Sqlmock) {
		rows := newMeetingRows()
		for _, expectedRow := range expectedMeetings {
			rows.AddRow(expectedRow.MeetingID, expectedRow.Name, expectedRow.Date, expectedRow.Started, expectedRow.CountConnected)
		}
		mock.ExpectQuery("SELECT").WithArgs(testUserID).WillReturnRows(rows).WillReturnError(nil)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		meetings := display.GetMyMeetings(testUserID)
		equalMeetings(t, expectedMeetings, meetings)
	}

	testingMockMeeting(t, initMock, testMock)
}

func equalMeetings(t *testing.T, expectedMeetings, meetings []meetingController.Meeting) {
	if len(meetings) != len(expectedMeetings) {
		t.Fatalf("expected %v, got %v", expectedMeetings, meetings)
	}
	for i := 0; i < len(meetings); i++ {
		if expectedMeetings[i].MeetingID != meetings[i].MeetingID ||
			expectedMeetings[i].Name != meetings[i].Name ||
			expectedMeetings[i].Date != meetings[i].Date ||
			expectedMeetings[i].Started != meetings[i].Started ||
			expectedMeetings[i].CountConnected != meetings[i].CountConnected {
			t.Fatalf("expected %v, got %v", expectedMeetings, meetings)
		}
	}
}

func TestStartMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectExec("UPDATE").WithArgs(testUserID, testMeetingID)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		display.StartMeeting(testUserID, testMeetingID)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestEndMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectExec("DELETE").WithArgs(testUserID, testMeetingID)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		display.EndMeeting(testUserID, testMeetingID)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestErrorGetCountConnectedConnectToMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.ConnectToMeeting(testMeetingID)
		checkErrorIsNotNil(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestErrorMostCountConnectedConnectToMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := newRowsCountConnected(12)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.ConnectToMeeting(testMeetingID)
		checkErrorIsNotNil(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func newRowsCountConnected(count int) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"count_connected"})
	rows.AddRow(count)
	return rows
}

func TestErrorUpdateCountConnectedConnectedConnectToMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := newRowsCountConnected(10)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)

		mock.ExpectExec("UPDATE").WithArgs(testMeetingID, 11).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.ConnectToMeeting(testMeetingID)
		checkErrorIsNotNil(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestSuccessfulConnectToMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := newRowsCountConnected(10)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)

		result := getEmptyResult()
		mock.ExpectExec("UPDATE").WithArgs(testMeetingID, 11).WillReturnResult(result)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		err := display.ConnectToMeeting(testMeetingID)
		checkErrorIsNil(t, err)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestDisconnectToMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := newRowsCountConnected(1)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)

		result := getEmptyResult()
		mock.ExpectExec("UPDATE").WithArgs(testMeetingID, 0).WillReturnResult(result)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		display.DisconnectToMeeting(testMeetingID)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestErrorFalseIsStartedMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		started := display.IsStartedMeeting(testMeetingID)
		checkIsFalse(t, started)
	}

	testingMockMeeting(t, initMock, testMock)
}

func TestSuccessfulFalseIsStartedMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := getRowsStarted(false)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		started := display.IsStartedMeeting(testMeetingID)
		checkIsFalse(t, started)
	}

	testingMockMeeting(t, initMock, testMock)
}

func getRowsStarted(value bool) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"started"})
	rows.AddRow(value)
	return rows
}

func TestSuccessfulTrueIsStartedMeeting(t *testing.T) {
	initMock := func(mock sqlmock.Sqlmock) {
		rows := getRowsStarted(true)
		mock.ExpectQuery("SELECT").WithArgs(testMeetingID).WillReturnRows(rows).WillReturnError(nil)
	}

	testMock := func(t *testing.T, display meetingController.Display) {
		started := display.IsStartedMeeting(testMeetingID)
		checkIsTrue(t, started)
	}

	testingMockMeeting(t, initMock, testMock)
}
