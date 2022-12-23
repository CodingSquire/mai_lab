package middlewares

import "net/http"

// Adapter is a function that adapts an http.Handler.
type Adapter func(http.Handler) http.Handler

// Adapt applies the given adapters to the given handler.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
