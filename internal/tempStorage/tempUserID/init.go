package tempUserID

type Display interface {
	TakeTempUserID() (any, error)
	GiveTempUserID(userID any)
}
