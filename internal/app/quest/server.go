package quest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rafaelzig/rest-quest/internal/pkg/log"
	"net/http"
	"time"
)

type Server struct {
	Router         *mux.Router
	LogHandlerFunc func(v interface{})
}

const supportedContentType = "application/json; charset=UTF-8"

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	header := w.Header()
	header.Set("Content-Type", supportedContentType)
	w.WriteHeader(status)
	if data == nil {
		return
	}
	err := s.encode(w, data)
	if err != nil {
		s.Error(fmt.Sprintf("respond failed: %s", err))
	}
}

func (s *Server) Trace(v interface{}) {
	s.log(log.TRACE, v)
}
func (s *Server) Debug(v interface{}) {
	s.log(log.DEBUG, v)
}
func (s *Server) Info(v interface{}) {
	s.log(log.INFO, v)
}
func (s *Server) Warn(v interface{}) {
	s.log(log.WARN, v)
}
func (s *Server) Error(v interface{}) {
	s.log(log.ERROR, v)
}
func (s *Server) Fatal(v interface{}) {
	s.log(log.FATAL, v)
}

func (s *Server) log(l log.Level, v interface{}) {
	if s.LogHandlerFunc == nil {
		return
	}
	e := &log.Event{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Level:     l,
		Message:   v,
	}
	s.LogHandlerFunc(e)
}
