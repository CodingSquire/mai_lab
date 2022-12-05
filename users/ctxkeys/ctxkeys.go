package ctxkeys

type contextKey string

var (
	ContextKeyParams = contextKey("params")
)
