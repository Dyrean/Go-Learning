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

	createTables(db)

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

func createTables(db *sql.DB) {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			owner_id TEXT NOT NULL
		)
	`

	if _, err := db.Query(createEventsTable); err != nil {
		log.Fatalf(fmt.Errorf("could not create events table: %w", err).Error())
	}
	log.Info("events table created")
}
