package quest

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Server) encode(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) decode(r *http.Request, v interface{}) error {
	if r.Header.Get("Content-Type") != supportedContentType {
		return errors.New("decode: Unsupported Content-Type")
	}
	return json.NewDecoder(r.Body).Decode(v)
}
