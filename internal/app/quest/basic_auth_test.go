package quest

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestBasicAuthPass(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	user := "user"
	pass := "pass"
	t.Cleanup(func() {
		os.Remove(pass)
	})
	_, err := os.Create(pass)
	is.NoErr(err)
	r.SetBasicAuth(user, pass)
	is.True(srv.basicAuth(user)(r))
}

func TestBasicAuthFail(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.SetBasicAuth("user", "pass")
	is.True(!srv.basicAuth("another")(r))
}
