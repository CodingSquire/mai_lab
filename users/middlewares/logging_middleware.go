// Package middlewares provides middlewares for the application.
package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// Message is a log message.
type Message struct {
	StatusCode int
	Method     string
	Path       string
	Time       time.Duration
}

// LoggingMiddleware is a middleware that logs requests.
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

// responseWriter wraps an http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

// Status returns the status code of the response.
func (rw *responseWriter) Status() int {
	return rw.status
}

// WriteHeader writes the status code to the response.
func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}
