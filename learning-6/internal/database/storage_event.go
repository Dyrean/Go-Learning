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

func (s *Service) GetEvents() (*[]Event, error) {
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

	return &events, nil
}

func (s *Service) SaveEvent(event Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	INSERT INTO events(id, name, description, date_time, owner_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, event.ID, event.Name, event.Description, event.DateTime, event.OwnerID, event.CreatedAt, event.UpdatedAt)
	if err != nil {
		return err
	}

	log.Infof("event added id: &v", event.ID)
	return nil
}

func (s *Service) GetEvent(id string) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	SELECT id, name, description, date_time, owner_id, created_at, updated_at
	FROM events
	WHERE id = ?
	`

	row := s.db.QueryRowContext(ctx, query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.OwnerID, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}

	log.Infof("get event id %v", event.ID)
	return &event, nil
}

func (s *Service) UpdateEvent(event Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	UPDATE events
	SET name = ?, description = ?, date_time = ?, owner_id = ?, updated_at = ?
	WHERE id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, event.Name, event.Description, event.DateTime, event.OwnerID, event.UpdatedAt, event.ID)
	if err != nil {
		return err
	}

	log.Infof("update event id: %v", event.ID)
	return nil
}

func (s *Service) DeleteEvent(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	DELETE FROM events
	WHERE id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	log.Infof("delete event db id: %v", id)
	return nil
}

func (s *Service) GetRegistrations(eventID string) (*[]Registeration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	SELECT id, event_id, user_id, created_at
	FROM registrations
	ORDER BY created_at DESC;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registerations []Registeration

	for rows.Next() {
		var registeration Registeration

		err := rows.Scan(&registeration.ID, &registeration.EventID, &registeration.UserID, &registeration.CreatedAt)
		if err != nil {
			return nil, err
		}
		registerations = append(registerations, registeration)
	}

	return &registerations, nil
}
