package server

func (s *FiberServer) RegisterRoutes() {
	s.App.Get("/health", s.healthHandler)
}
