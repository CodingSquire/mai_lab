package ctxkeys

type contextKey string

var (
	// ContextKeyParams is the key for the request parameters.
	ContextKeyParams = contextKey("params")
)
