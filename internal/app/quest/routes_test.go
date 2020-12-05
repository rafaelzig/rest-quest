package quest

import (
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesAllowedMethodsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	is := is.New(t)
	srv := Server{Router: mux.NewRouter()}
	srv.Routes()
	routes := map[string]string{
		"/":        http.MethodGet,
		"/started": http.MethodGet,
		"/health":  http.MethodGet,
	}

	for route, method := range routes {
		r := httptest.NewRequest(method, route, nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, r)
		is.True(w.Code != http.StatusNotFound)
		is.True(w.Code != http.StatusMethodNotAllowed)
	}
}

func TestRoutesDisallowedMethodsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	is := is.New(t)
	srv := Server{Router: mux.NewRouter()}
	srv.Routes()
	routes := map[string][8]string{
		"/": {
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
			http.MethodDelete,
		},
		"/health": {
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
			http.MethodDelete,
		},
	}
	for route, disallowedMethods := range routes {
		for _, disallowedMethod := range disallowedMethods {
			r := httptest.NewRequest(disallowedMethod, route, nil)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, r)
			is.Equal(w.Code, http.StatusMethodNotAllowed)
		}
	}
}
