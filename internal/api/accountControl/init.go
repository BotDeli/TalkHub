package accountControl

import (
	"TalkHub/internal/api/accountControl/pb"
	"TalkHub/internal/config"
	"log"
)

type Display interface {
	Registration(email, password string) (session *pb.SessionData, strErr string)
	Authorization(email, password string) (session *pb.SessionData, strErr string)
	ChangePassword(email, password, newPassword string) (strErr string)
	DeleteAccount(id, email, password string)
	IsAuthenticated(key string) (id string, strErr string)
	DeleteSession(key string)
}

func MustInitDisplay(cfg *config.GRPCConfig) Display {
	displayG, err := initGRPCClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return displayG
}
