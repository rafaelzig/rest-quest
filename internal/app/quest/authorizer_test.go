package quest

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckAuthorization(t *testing.T) {
	is := is.New(t)
	var isCalled bool
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isCalled = true
	})
	authFunc := func(r *http.Request) bool {
		return false
	}
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	srv.WithAuth(authFunc, h)(w, r)
	is.Equal(w.Code, http.StatusUnauthorized)
	is.Equal(isCalled, false)

	authFunc = func(r *http.Request) bool {
		return true
	}
	srv.WithAuth(authFunc, h)(w, r)
	is.Equal(isCalled, true)
}
