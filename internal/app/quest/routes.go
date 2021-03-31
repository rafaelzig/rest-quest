package quest

import "net/http"

func (s *Server) Routes() {
	s.Router.Use(s.LoggingMiddleware())
	s.Router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
	s.Router.HandleFunc("/started", s.handleStarted()).Methods(http.MethodGet).Queries("name", "{name:[a-zA-Z]{1,10}}")
	s.Router.HandleFunc("/started", s.handleStartedHelp()).Methods(http.MethodGet)
	s.Router.HandleFunc("/north", s.WithAuth(s.basicAuth("warrior"), s.handleNorth())).Methods(http.MethodGet)
	s.Router.HandleFunc("/health", s.handleHealth()).Methods(http.MethodGet)
}
