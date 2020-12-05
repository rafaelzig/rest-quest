package quest

import (
	"net/http"
	"time"
)

type responseWriter struct {
	status int
	body   string
	http.ResponseWriter
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(body []byte) (int, error) {
	w.body = string(body)
	return w.ResponseWriter.Write(body)
}

func (s *Server) LoggingMiddleware() func(http.Handler) http.Handler {
	type request struct {
		Method string `json:"method"`
		URI    string `json:"uri"`
	}
	type response struct {
		Status  int           `json:"status"`
		Body    string        `json:"body,omitempty"`
		Elapsed time.Duration `json:"elapsed"`
	}
	type event struct {
		Request  request  `json:"request"`
		Response response `json:"response"`
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lw := &responseWriter{
				ResponseWriter: w,
			}
			h.ServeHTTP(lw, r)
			elapsed := time.Since(start)
			e := &event{
				Request: request{
					Method: r.Method,
					URI:    r.RequestURI,
				},
				Response: response{
					Status:  lw.status,
					Body:    lw.body,
					Elapsed: elapsed,
				},
			}
			s.Info(e)
		})
	}
}
