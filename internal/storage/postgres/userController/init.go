package userController

type Display interface {
	SaveUserInfo(u *User)
	GetUserInfoFromEmail(email string) (*User, error)
	GetUserInfoFromID(id any) (*User, error)
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
