package userController

import (
	"TalkHub/internal/storage/postgres"
	"database/sql"
	"log"
)

type UIDisplay struct {
	*postgres.Storage
}

func InitDisplay(pg *postgres.Storage) Display {
	initTable(pg.DB)
	return &UIDisplay{Storage: pg}
}

func initTable(db *sql.DB) {
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    	id VARCHAR NOT NULL PRIMARY KEY UNIQUE,
		first_name VARCHAR NOT NULL,
		last_name VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE
	)`); err != nil {
		log.Printf("Error creating users table: %s\n", err)
	}
}

func (uid *UIDisplay) SaveUserInfo(u *User) {
	query := `INSERT INTO users (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := uid.DB.Exec(
		query,
		u.Id,
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
	rows, err := uid.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	return scanUserInfo(rows)
}

func scanUserInfo(rows *sql.Rows) (*User, error) {
	var u User
	rows.Next()
	err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (uid *UIDisplay) GetUserInfoFromID(id any) (*User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	rows, err := uid.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	return scanUserInfo(rows)
}
