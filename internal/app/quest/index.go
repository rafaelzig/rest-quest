package quest

import "net/http"

func (s *Server) handleIndex() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Message string      `json:"message"`
		Next    interface{} `json:"next"`
	}{
		Message: "Hello young warrior, please proceed with caution",
		Next: struct {
			Action   string `json:"action"`
			Location string `json:"location"`
		}{
			Action:   http.MethodGet,
			Location: "/started",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}
