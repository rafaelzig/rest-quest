package quest

import (
	"encoding/json"
	"fmt"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleNorth(t *testing.T) {
	is := is.New(t)
	name := "stranger"
	t.Cleanup(func() {
		os.Remove(name)
	})
	_, err := os.Create(name)
	is.NoErr(err)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/north", name), nil)
	w := httptest.NewRecorder()
	srv.handleStarted()(w, r)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Header().Get("Content-Type"), supportedContentType)

	type next = struct {
		Action   string `json:"action"`
		Location string `json:"location"`
	}
	type response = struct {
		Message string `json:"message"`
		Next    next   `json:"next"`
	}
	expected := response{
		Message: fmt.Sprintf("You have returned, brave warrior %s! You seem lost.", name),
		Next: next{
			Action:   http.MethodHead,
			Location: "/north",
		},
	}

	var actual response
	err = json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
