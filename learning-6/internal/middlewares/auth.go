package middlewares

import (
	"event-booking/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Authenticate(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		log.Warn("no token provided")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Not authorized"})
	}

	token = token[len("Bearer "):]

	userID, err := utils.ValidateToken(token)
	if err != nil {
		log.Warnf("could not validate token: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Not authorized"})
	}
	c.Locals("userID", userID)
	return c.Next()
}
