package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

type Service struct {
	db *sql.DB
}

const MaxOpenConns = 100
const MaxIdleConns = 10

var (
	dburl = os.Getenv("DB_URL")
)

func New() *Service {
	db, err := sql.Open("sqlite", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	createDBSchema(db)

	return &Service{db: db}
}

func (s *Service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Errorf("db down: %w", err).Error())
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func createDBSchema(db *sql.DB) {
	schema := `
		PRAGMA foreign_keys = ON;

		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);

		CREATE TABLE IF NOT EXISTS events (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			owner_id TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS registrations (
			id TEXT PRIMARY KEY,
			event_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
			UNIQUE (event_id, user_id)
		);
	`
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf(fmt.Errorf("could not create schema: %w", err).Error())
	}
}
