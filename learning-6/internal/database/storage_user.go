package database

import (
	"context"
	"errors"
	"event-booking/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `binding:"required" json:"email"`
	Password  string    `binding:"required" json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (s *Service) SaveUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	INSERT INTO users(id, email, password, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.ID, user.Email, user.Password, user.CreatedAt, user.CreatedAt)
	if err != nil {
		return err
	}

	total, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Infof("signup user db row affected: %v, email: %v", total, user.Email)
	return nil
}

func (s *Service) ValidateCredentials(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT password FROM users WHERE email = ?`
	row := s.db.QueryRowContext(ctx, query, user.Email)

	var password string
	err := row.Scan(&password)
	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(user.Password, password)

	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
