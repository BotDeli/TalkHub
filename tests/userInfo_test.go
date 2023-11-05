package tests

import (
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/userInfo"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

var (
	testUser = &userInfo.User{
		UserIcon:  "",
		FirstName: "TestFirstName",
		LastName:  "T E S T L A S T N A M E",
		Email:     "Email@example.com",
	}

	testError = errors.New("test error")

	tableColumns = []string{"user_icon", "first_name", "last_name", "email"}
)

func testingMockUser(t *testing.T, initMock func(sqlmock.Sqlmock), testMock func(*testing.T, userInfo.Display)) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	initMock(mockDB)

	displayU := &userInfo.UIDisplay{PG: &postgres.Storage{DB: db}}

	testMock(t, displayU)
}

func TestSuccessfulSaveUserInfo(t *testing.T) {
	emptyResult := getEmptyResult()

	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectExec("INSERT INTO").WithArgs(
			testUser.UserIcon,
			testUser.FirstName,
			testUser.LastName,
			testUser.Email,
		).WillReturnResult(emptyResult)
	}

	testMock := func(t *testing.T, displayU userInfo.Display) {
		displayU.SaveUserInfo(testUser)
	}

	testingMockUser(t, initMock, testMock)
}

func getEmptyResult() driver.Result {
	return sqlmock.NewResult(0, 0)
}

func TestErrorGetUserInfo(t *testing.T) {
	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnError(testError)
	}

	testMock := func(t *testing.T, displayU userInfo.Display) {
		u, err := displayU.GetUserInfo(testUser.Email)
		checkUserIsNil(t, u)
		checkErrorIsTestError(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func checkUserIsNil(t *testing.T, u *userInfo.User) {
	if u != nil {
		t.Error("expected user info is nil, got ", u)
	}
}

func checkErrorIsTestError(t *testing.T, err error) {
	if err != testError {
		t.Error("expected test error, got", err)
	}
}

func TestEmptyUserInfoGetUserInfo(t *testing.T) {
	emptyRows := getEmptyRows()

	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnRows(emptyRows)
	}

	testMock := func(t *testing.T, displayU userInfo.Display) {
		u, err := displayU.GetUserInfo(testUser.Email)
		checkUserIsNil(t, u)
		checkErrorIsNotNil(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func getEmptyRows() *sqlmock.Rows {
	return sqlmock.NewRows(tableColumns)
}

func checkErrorIsNotNil(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error, got", err)
	}
}

func TestSuccessfulGetUserInfo(t *testing.T) {
	rows := getEmptyRows()
	rows.AddRow(
		testUser.UserIcon,
		testUser.FirstName,
		testUser.LastName,
		testUser.Email,
	)

	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnRows(rows)
	}

	testMock := func(t *testing.T, displayU userInfo.Display) {
		u, err := displayU.GetUserInfo(testUser.Email)
		if u == nil {
			t.Fatal("user info is nil")
		}
		equalsUserInfo(t, testUser, u)
		checkErrorIsNil(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func equalsUserInfo(t *testing.T, u1, u2 *userInfo.User) {
	if u1.UserIcon != u2.UserIcon {
		t.Errorf("u1 user icon %s != u2 user icon %s", u1.UserIcon, u2.UserIcon)
	}
	if u1.FirstName != u2.FirstName {
		t.Errorf("u1 first name %s != u2 first name %s", u1.FirstName, u2.FirstName)
	}
	if u1.LastName != u2.LastName {
		t.Errorf("u1 last name %s != u2 last name %s", u1.LastName, u2.LastName)
	}
	if u1.Email != u2.Email {
		t.Errorf("u1 email %s != u2 email %s", u1.Email, u2.Email)
	}
}

func checkErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Error("expected nil, got", err)
	}
}
