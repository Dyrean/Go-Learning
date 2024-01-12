package database

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Service) GetEvents() ([]Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	SELECT id, name, description, date_time, owner_id, created_at, updated_at
	FROM events
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.OwnerID, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (s *Service) SaveEvent(event Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	INSERT INTO events(id, name, description, date_time, owner_id, created_at)
	VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, event.ID, event.Name, event.Description, event.DateTime, event.OwnerID, event.CreatedAt)
	if err != nil {
		return err
	}

	total, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Infof("events db row affected: %v", total)
	return nil
}
