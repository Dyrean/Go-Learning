package server

import (
	"database/sql"
	"event-booking/internal/database"
	"event-booking/internal/utils"
	"fmt"
	"strings"
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
		log.Warnf("could not parse event save request data: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Errorf("could not parse event save request data: %w", err).Error()})
	}

	if time.Now().After(event.DateTime) {
		log.Warn("event end date is in the past")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "event end date is in the past"})
	}

	jwtUserID := c.Locals("userID").(string)

	event.ID = uuid.NewString()
	event.CreatedAt = time.Now()
	event.UpdatedAt = event.CreatedAt
	event.OwnerID = jwtUserID

	err := s.db.SaveEvent(event)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "foreign key constraint") || strings.Contains(strings.ToLower(err.Error()), "violates foreign key") {
			log.Warnf("bad request could not create event: %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not create event."})
		}
		log.Warnf("could not create event: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not create event."})
	}
	return c.Status(fiber.StatusCreated).JSON(event)
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

func (s *FiberServer) updateEvent(c *fiber.Ctx) error {
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

	jwtUserID := c.Locals("userID").(string)

	if event.OwnerID != jwtUserID {
		log.Warnf("user with id: %s is not the owner of the event to update", jwtUserID)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Not authorized to update the event"})
	}

	var updatedEvent database.Event
	if err := c.BodyParser(&updatedEvent); err != nil {
		log.Warnf("could not parse update event request data: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Errorf("could not parse update event request data: %w", err).Error()})
	}

	updatedEvent.ID = id
	updatedEvent.UpdatedAt = time.Now()
	updatedEvent.CreatedAt = event.CreatedAt
	updatedEvent.OwnerID = event.OwnerID

	log.Infof("updated event: %+v", updatedEvent)

	if updatedEvent.DateTime.IsZero() {
		updatedEvent.DateTime = event.DateTime
	}
	if updatedEvent.Name == "" {
		updatedEvent.Name = event.Name
	}
	if updatedEvent.Description == "" {
		updatedEvent.Description = event.Description
	}
	if updatedEvent.OwnerID == "" {
		updatedEvent.OwnerID = event.OwnerID
	}

	if time.Now().After(updatedEvent.DateTime) {
		log.Warn("event end date is in the past")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "event end date is in the past"})
	}

	err = s.db.UpdateEvent(updatedEvent)
	if err != nil {
		log.Warnf("could not update event: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not update event."})
	}
	return c.JSON(fiber.Map{"message": "event updated"})
}

func (s *FiberServer) deleteEvent(c *fiber.Ctx) error {
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

	jwtUserID := c.Locals("userID").(string)

	if event.OwnerID != jwtUserID {
		log.Warnf("user with id: %s is not the owner of the event to delete", jwtUserID)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Not authorized to delete the event"})
	}

	err = s.db.DeleteEvent(id)
	if err != nil {
		log.Warnf("could not delete event: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not delete event."})
	}

	return c.JSON(fiber.Map{"message": "event deleted"})
}

func (s *FiberServer) signUp(c *fiber.Ctx) error {
	var user database.User
	if err := c.BodyParser(&user); err != nil {
		log.Warnf("could not parse signup request data: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Errorf("could not parse signup request data: %w", err).Error()})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Warnf("could not hash password: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not signup user."})
	}

	user.ID = uuid.NewString()
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	err = s.db.SaveUser(user)
	if err != nil {
		log.Warnf("could not signup user: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not signup user."})
	}

	return c.JSON(fiber.Map{"message": "signup successful"})
}

func (s *FiberServer) login(c *fiber.Ctx) error {
	var user database.User
	if err := c.BodyParser(&user); err != nil {
		log.Warnf("could not parse login request data: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Errorf("could not parse login request data: %w", err).Error()})
	}

	if user.Email == "" || user.Password == "" {
		log.Warn("email or password is empty")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "email or password is empty"})
	}

	err := s.db.ValidateCredentials(&user)
	if err != nil {
		log.Warnf("Validate error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Warnf("generate token error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "credentials invalid"})
	}

	return c.JSON(fiber.Map{"message": "login successful", "token": token})
}
