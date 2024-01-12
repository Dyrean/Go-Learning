package server

import (
	"database/sql"
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
	events, err := s.db.GetEvents()
	if err != nil {
		log.Warnf("could not fetch events: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not fetch events."})
	}

	return c.JSON(fiber.Map{"events": events})
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

	err := s.db.SaveEvent(event)
	if err != nil {
		log.Warnf("could not create event: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not create event."})
	}
	return c.JSON(event)
}

func (s *FiberServer) getEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	event, err := s.db.GetEvent(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("no event with id: %s", id)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("no event with id: %v", id)})
		}
		log.Warnf("could not fetch event: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not fetch event."})
	}

	return c.JSON(event)
}
