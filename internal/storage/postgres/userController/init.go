package userController

type Display interface {
	SaveUserInfo(u *User)
	GetUserInfo(email string) (*User, error)
}

type User struct {
	Id        string
	UserIcon  string
	FirstName string
	LastName  string
	Email     string
}
