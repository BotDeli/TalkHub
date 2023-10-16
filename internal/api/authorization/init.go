package authorization

import (
	"TalkHub/internal/config"
	"log"
)

type Display interface {
	Register(login, password string) (key string, strErr string)
	LogIn(login, password string) (key string, strErr string)
	IsAuthenticated(key string) (login string, strErr string)
	ChangePassword(login, password, newPassword string) (strErr string)
}

func MustInitDisplay(cfg *config.GRPCConfig) Display {
	display, err := initGRPCClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return display
}
