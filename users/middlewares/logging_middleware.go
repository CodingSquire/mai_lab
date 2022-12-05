package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type Message struct {
	StatusCode int
	Method     string
	Path       string
	Time       time.Duration
}

func LoggingMiddleware() Adapter {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Print(
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			message := Message{
				StatusCode: wrapped.status,
				Method:     r.Method,
				Path:       r.URL.EscapedPath(),
				Time:       time.Since(start),
			}
			json.MarshalIndent(message, "", "  ")
			log.Print(message)
		}
		return http.HandlerFunc(fn)
	}
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}
