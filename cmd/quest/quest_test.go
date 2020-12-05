package main_test

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"github.com/rafaelzig/rest-quest/internal/app/quest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	is := is.New(t)
	h := &quest.Server{Router: mux.NewRouter()}
	h.Routes()
	srv := httptest.NewServer(h)
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/health")
	is.NoErr(err)
	is.Equal(resp.StatusCode, http.StatusOK)
	is.Equal(resp.Header.Get("Content-Type"), "application/json; charset=UTF-8")
	type response = struct {
		Status string `json:"status"`
	}
	expected := response{
		Status: "ready",
	}
	var actual response
	err = json.NewDecoder(resp.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
