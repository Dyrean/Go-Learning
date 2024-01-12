package database

import (
	"context"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}

var events = []Event{}

func (s *Service) GetEvents() []Event {
	_, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	return events
}

func (s *Service) SaveEvent(event Event) error {
	events = append(events, event)
	return nil
}
