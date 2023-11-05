package postgres

import (
	"TalkHub/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	DB *sql.DB
}

func InitPostgres(cfg *config.PostgresConfig) (*Storage, error) {
	db, err := sql.Open("postgres", cfg.GetSourceName())
	if err != nil {
		log.Println("Error connection to postgres", cfg.GetSourceName())
		return nil, err
	}

	return &Storage{DB: db}, db.Ping()
}

func (pg *Storage) Close() {
	if err := pg.DB.Close(); err != nil {
		log.Println("Error close postgres", err)
	}
}
