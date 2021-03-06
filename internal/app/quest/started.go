package quest

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) handleStartedHelp() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Message string      `json:"message"`
		Next    interface{} `json:"next"`
	}{
		Message: "An unnamed warrior is unwelcome!",
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
		s.respond(w, r, response, http.StatusUnauthorized)
	}
}

func (s *Server) handleStarted() func(http.ResponseWriter, *http.Request) {
	type next struct {
		Action   string `json:"action"`
		Location string `json:"location"`
	}
	type response struct {
		Message string `json:"message"`
		Next    next   `json:"next"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		exists, err := fileExists(name)
		if err != nil {
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		if exists {
			resp := response{
				Message: fmt.Sprintf("You have returned, brave warrior %s! You seem lost.", name),
				Next: next{
					Action:   http.MethodHead,
					Location: "/north",
				},
			}
			s.respond(w, r, resp, http.StatusOK)
			return
		}
		_, err = os.Create(name)
		if err != nil {
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		resp := response{
			Message: fmt.Sprintf("Welcome, brave warrior %s! Head north from here to where your quest awaits you.", name),
			Next: next{
				Action:   http.MethodHead,
				Location: "/north",
			},
		}
		s.respond(w, r, resp, http.StatusCreated)
	}
}
