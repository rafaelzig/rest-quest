package quest

import (
	"net/http"
	"os"
)

func (s *Server) basicAuth(u string) func(*http.Request) bool {
	return func(r *http.Request) bool {
		username, password, ok := r.BasicAuth()
		if !ok {
			return false
		}
		if _, err := os.Open(password); err != nil {
			return false
		}
		return u == username
	}
}
