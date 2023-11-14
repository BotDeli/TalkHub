package userController

import (
	"TalkHub/internal/storage/postgres"
	"database/sql"
	"fmt"
	"log"
)

type UIDisplay struct {
	PG *postgres.Storage
}

func InitDisplay(pg *postgres.Storage) Display {
	initTable(pg.DB)
	return &UIDisplay{PG: pg}
}

func initTable(db *sql.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    	id VARCHAR NOT NULL PRIMARY KEY UNIQUE,
    	user_icon BYTEA,
		first_name VARCHAR NOT NULL,
		last_name VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE
	)`); err != nil {
		log.Printf("Error creating users table: %s\n", err)
	}
}

func (uid *UIDisplay) SaveUserInfo(u *User) {
	query := `INSERT INTO users (id, user_icon, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5)`
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

func (uid *UIDisplay) GetUserInfoFromEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	rows, err := uid.PG.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	return scanUserInfo(rows)
}

func scanUserInfo(rows *sql.Rows) (*User, error) {
	var u User
	rows.Next()
	err := rows.Scan(&u.Id, &u.UserIcon, &u.FirstName, &u.LastName, &u.Email)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (uid *UIDisplay) GetUserInfoFromID(id any) (*User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	rows, err := uid.PG.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	return scanUserInfo(rows)
}
