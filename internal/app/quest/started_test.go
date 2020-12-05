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

func TestHandleStartedHelpResponse(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/started", nil)
	w := httptest.NewRecorder()
	srv.handleStartedHelp()(w, r)
	is.Equal(w.Code, http.StatusUnauthorized)
	is.Equal(w.Header().Get("Content-Type"), supportedContentType)

	type next = struct {
		Action      string `json:"action"`
		Location    string `json:"location"`
		Requirement string `json:"requirement"`
	}
	type response = struct {
		Message string `json:"message"`
		Next    next   `json:"next"`
	}
	expected := response{
		Message: "An unnamed warrior is unwelcome!",
		Next: next{
			Action:      http.MethodGet,
			Location:    "/started",
			Requirement: "name",
		},
	}

	var actual response
	err := json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}

func TestHandleStartedNewUserResponse(t *testing.T) {
	name := "stranger"
	t.Cleanup(func() {
		os.Remove(name)
	})
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/started?name=%s", name), nil)
	w := httptest.NewRecorder()
	srv.handleStarted()(w, r)
	is.Equal(w.Code, http.StatusCreated)
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
		Message: fmt.Sprintf("Welcome, brave warrior %s! Head north from here to where your quest awaits you.", name),
		Next: next{
			Action:   http.MethodHead,
			Location: "/north",
		},
	}

	var actual response
	err := json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}

func TestHandleStartedExistingUserResponse(t *testing.T) {
	is := is.New(t)
	name := "stranger"
	t.Cleanup(func() {
		os.Remove(name)
	})
	_, err := os.Create(name)
	is.NoErr(err)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/started?name=%s", name), nil)
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
