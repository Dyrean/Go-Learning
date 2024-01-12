package server

import (
	"event-booking/internal/database"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) getEvents(c *fiber.Ctx) error {
	events := s.db.GetEvents()
	log.Infof("events: %v", events)
	return c.JSON(events)
}

func (s *FiberServer) saveEvent(c *fiber.Ctx) error {
	var event database.Event
	if err := c.BodyParser(&event); err != nil {
		log.Warnf("could not parse request data: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Errorf("could not parse request data: %w", err).Error()})
	}

	if time.Now().After(event.DateTime) {
		log.Warn("event end date is in the past")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "event end date is in the past"})
	}

	event.ID = uuid.NewString()
	event.OwnerID = uuid.NewString()

	event.CreatedAt = time.Now()

	s.db.SaveEvent(event)
	return c.JSON(event)
}
