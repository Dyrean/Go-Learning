package database

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Registeration struct {
	ID        string    `binding:"required" json:"id"`
	UserID    string    `binding:"required" json:"user_id"`
	EventID   string    `binding:"required" json:"event_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Service) CreateRegistration(registeration *Registeration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	INSERT INTO registrations(id, event_id, user_id, created_at)
	VALUES (?, ?, ?, ?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, registeration.ID, registeration.EventID, registeration.UserID, registeration.CreatedAt)
	if err != nil {
		return err
	}

	log.Infof("registeration added id: &v", registeration.ID)
	return nil
}

func (s *Service) GetRegistration(userID, eventID string) (*Registeration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	SELECT id, event_id, user_id, created_at
	FROM registrations
	WHERE user_id = ? AND event_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var registeration Registeration
	err = stmt.QueryRowContext(ctx, userID, eventID).Scan(
		&registeration.ID,
		&registeration.EventID,
		&registeration.UserID,
		&registeration.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &registeration, nil
}

func (s *Service) CancelRegisteration(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	DELETE FROM registrations
	WHERE id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	log.Infof("cancel registeration id: %v", id)
	return nil
}
