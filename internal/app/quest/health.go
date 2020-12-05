package quest

import (
	"net/http"
)

func (s *Server) handleHealth() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "ready",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}
