package server

import "event-booking/internal/middlewares"

func (s *FiberServer) RegisterRoutes() {
	s.App.Get("/health", s.healthHandler)

	s.RegisterEventRoutes()
}

func (s *FiberServer) RegisterEventRoutes() {
	s.App.Get("/events", s.getEvents)
	s.App.Get("/events/:id", s.getEvent)

	s.App.Post("/events", middlewares.Authenticate, s.saveEvent)
	s.App.Put("/events/:id", s.updateEvent)
	s.App.Delete("/events/:id", s.deleteEvent)

	s.App.Post("/signup", s.signUp)
	s.App.Post("/login", s.login)
}
