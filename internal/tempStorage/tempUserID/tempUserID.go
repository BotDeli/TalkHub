package tempUserID

import (
	"TalkHub/internal/storage/postgres/userController"
	"TalkHub/pkg/generator"
	"errors"
)

var errExceededCountAttemptsCreateTempUserID = errors.New("превышено количество попыток создать временный ид")

type StorageTempUserID struct {
	storage  map[any]bool
	displayU userController.Display
}

func InitDisplay(displayU userController.Display) Display {
	return &StorageTempUserID{
		storage:  make(map[any]bool),
		displayU: displayU,
	}
}

func (s *StorageTempUserID) TakeTempUserID() (any, error) {
	var userID string
	for i := 0; i < 5; i++ {
		userID = generator.NewUUIDDigits()

		if s.storage[userID] {
			continue
		}

		_, err := s.displayU.GetUserInfoFromID(userID)
		if err != nil {
			s.storage[userID] = true
			return userID, nil
		}
	}
	return "", errExceededCountAttemptsCreateTempUserID
}

func (s *StorageTempUserID) GiveTempUserID(userID any) {
	delete(s.storage, userID)
}
