package userController

type Display interface {
	SaveUserInfo(u *User)
	GetUserInfoFromEmail(email string) (*User, error)
	GetUserInfoFromID(id any) (*User, error)
}

type User struct {
	Id        string
	UserIcon  string
	FirstName string
	LastName  string
	Email     string
}
