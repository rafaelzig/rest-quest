package quest

import (
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncode(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	w := httptest.NewRecorder()
	type response = struct {
		Message string `json:"message"`
	}
	expected := response{
		Message: "Secret World",
	}
	err := srv.encode(w, expected)
	is.NoErr(err)
	var actual response
	err = json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(expected, actual)
}

func TestDecode(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	type response = struct {
		Message string `json:"message"`
	}
	expected := response{
		Message: "Secret World",
	}
	w := httptest.NewRecorder()
	err := json.NewEncoder(w).Encode(expected)
	is.NoErr(err)
	r := httptest.NewRequest(http.MethodGet, "/", w.Body)
	r.Header.Set("Content-Type", supportedContentType)
	var actual response
	err = srv.decode(r, &actual)
	is.NoErr(err)
	is.Equal(expected, actual)
}
