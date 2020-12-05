package quest

import (
	"net/http"
)

func (s *Server) WithAuth(authFunc func(r *http.Request) bool, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if authFunc(r) {
			h(w, r)
		} else {
			s.respond(w, r, nil, http.StatusUnauthorized)
		}
	}
}
