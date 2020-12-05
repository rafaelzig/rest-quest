package quest

import "net/http"

func (s *Server) Routes() {
	s.Router.Use(s.LoggingMiddleware())
	s.Router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
	s.Router.HandleFunc("/health", s.handleHealth()).Methods(http.MethodGet)
}
