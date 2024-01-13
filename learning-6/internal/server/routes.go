package server

import "event-booking/internal/middlewares"

func (s *FiberServer) RegisterRoutes() {
	s.App.Get("/health", s.healthHandler)

	s.RegisterEventRoutes()
}

func (s *FiberServer) RegisterEventRoutes() {
	s.App.Get("/events", s.getEvents)
	s.App.Get("/events/:id", s.getEvent)

	s.App.Post("/signup", s.signUp)
	s.App.Post("/login", s.login)

	auth := s.App.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.Post("/events", s.saveEvent)
	auth.Put("/events/:id", s.updateEvent)
	auth.Delete("/events/:id", s.deleteEvent)
}
