package tests

import (
	"TalkHub/internal/storage/postgres"
	"TalkHub/internal/storage/postgres/userController"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

var (
	testUser = &userController.User{
		UserIcon:  "",
		FirstName: "TestFirstName",
		LastName:  "T E S T L A S T N A M E",
		Email:     "Email@example.com",
	}
)

func testingMockUser(t *testing.T, initMock func(sqlmock.Sqlmock), testMock func(*testing.T, userController.Display)) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	initMock(mockDB)

	display := &userController.UIDisplay{Storage: &postgres.Storage{DB: db}}

	testMock(t, display)
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

	testMock := func(t *testing.T, display userController.Display) {
		display.SaveUserInfo(testUser)
	}

	testingMockUser(t, initMock, testMock)
}

func TestErrorGetUserInfoFromEmail(t *testing.T) {
	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnError(testError)
	}

	testMock := func(t *testing.T, display userController.Display) {
		u, err := display.GetUserInfoFromEmail(testUser.Email)
		checkUserIsNil(t, u)
		checkErrorIsNotNil(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func checkUserIsNil(t *testing.T, u *userController.User) {
	if u != nil {
		t.Error("expected user info is nil, got ", u)
	}
}

func TestEmptyUserInfoGetUserInfoFromEmail(t *testing.T) {
	emptyRows := getEmptyUserRows()

	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnRows(emptyRows)
	}

	testMock := func(t *testing.T, display userController.Display) {
		u, err := display.GetUserInfoFromEmail(testUser.Email)
		checkUserIsNil(t, u)
		checkErrorIsNotNil(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func getEmptyUserRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "user_icon", "first_name", "last_name", "email"})
}

func TestSuccessfulGetUserInfoFromEmail(t *testing.T) {
	rows := getEmptyUserRows()
	rows.AddRow(
		testUser.Id,
		testUser.UserIcon,
		testUser.FirstName,
		testUser.LastName,
		testUser.Email,
	)

	initMock := func(mockDB sqlmock.Sqlmock) {
		mockDB.ExpectQuery("SELECT").WithArgs(testUser.Email).WillReturnRows(rows)
	}

	testMock := func(t *testing.T, display userController.Display) {
		u, err := display.GetUserInfoFromEmail(testUser.Email)
		if u == nil {
			t.Fatal("user info is nil")
		}
		equalsUserInfo(t, testUser, u)
		checkErrorIsNil(t, err)
	}

	testingMockUser(t, initMock, testMock)
}

func equalsUserInfo(t *testing.T, u1, u2 *userController.User) {
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
