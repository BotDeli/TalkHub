package userInfo

import (
	"TalkHub/internal/storage/postgres"
	"database/sql"
	"log"
)

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

type UIDisplay struct {
	PG *postgres.Storage
}

func InitDisplay(pg *postgres.Storage) Display {
	initTable(pg.DB)
	return &UIDisplay{PG: pg}
}

func initTable(db *sql.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users_info (
    	id VARCHAR NOT NULL PRIMARY KEY UNIQUE,
    	user_icon BYTEA,
		first_name VARCHAR NOT NULL,
		last_name VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE
	)`); err != nil {
		log.Printf("Error creating tables: %s\n", err)
	}
}

func (uid *UIDisplay) SaveUserInfo(u *User) {
	query := `INSERT INTO users_info (id, user_icon, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5)`
	_, err := uid.PG.DB.Exec(
		query,
		u.Id,
		u.UserIcon,
		u.FirstName,
		u.LastName,
		u.Email,
	)
	if err != nil {
		log.Println("SaveUserInfo error:", err)
	}
}

func (uid *UIDisplay) GetUserInfo(email string) (*User, error) {
	query := `SELECT * FROM users_info WHERE email = $1`
	var u User
	rows, err := uid.PG.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	rows.Next()
	err = rows.Scan(&u.Id, &u.UserIcon, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
