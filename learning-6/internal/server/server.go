package server

import (
	"event-booking/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServer struct {
	*fiber.App
	db *database.Service
}

func New() *FiberServer {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/metrics", monitor.New())

	return &FiberServer{
		App: app,
		db:  database.New(),
	}
}
