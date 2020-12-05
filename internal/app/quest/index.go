package quest

import "net/http"

func (s *Server) handleIndex() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Message string      `json:"message"`
		Next    interface{} `json:"next"`
	}{
		Message: "Hello brave warrior. In order to get started, a warrior needs a name!",
		Next: struct {
			Action      string `json:"action"`
			Location    string `json:"location"`
			Requirement string `json:"requirement"`
		}{
			Action:      http.MethodGet,
			Location:    "/started",
			Requirement: "name",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}
