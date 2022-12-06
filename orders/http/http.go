package http

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type HttpApp struct {
	router Router
	states map[string]interface{}
}

func NewApp() *HttpApp {
	return &HttpApp{
		router: Router{},
		states: make(map[string]interface{}),
	}
}

func (h *HttpApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// setting knows state
	for key, state := range h.states {
		ctx := context.WithValue(r.Context(), key, state)
		r = r.WithContext(ctx)
	}

	h.router.RunPreMiddleware(w, r)
	h.router.ServeHTTP(w, r)
}

func (a *HttpApp) Run(port string) error {
	a.router.PrettyPrint()
	fmt.Printf("Listening on port: %s\n", port)
	return http.ListenAndServe(":"+port, a)
}

func parsePattern(pattern string) *regexp.Regexp {
	splited := strings.Split(pattern, "/")
	buffer := new(strings.Builder)

	for _, canBePattern := range splited {
		if canBePattern != "" {
			buffer.WriteString(`\/`)
			if strings.HasPrefix(canBePattern, ":") {
				buffer.WriteString(`(?P<`)
				buffer.WriteString(canBePattern[1:])
				buffer.WriteString(`>.+)`)
			} else {
				buffer.WriteString(canBePattern)
			}
		}
	}

	return regexp.MustCompile(buffer.String())
}

func (a *HttpApp) Manage(key string, state interface{}) {
	a.states[key] = state
}

func (a *HttpApp) SetHandler(method string, strPattern string, handler RouterHandler) {
	pattern := parsePattern(strPattern)

	newroutes := append(a.router.routes, innerRouter{
		origPath: strPattern,
		method:   method,
		pattern:  pattern,
		handler:  handler,
	})

	a.router.routes = newroutes
}

func (a *HttpApp) Get(pattern string, handler RouterHandler) {
	a.SetHandler(http.MethodGet, pattern, handler)
}
func (a *HttpApp) Post(pattern string, handler RouterHandler) {
	a.SetHandler(http.MethodPost, pattern, handler)
}
func (a *HttpApp) Delete(pattern string, handler RouterHandler) {
	a.SetHandler(http.MethodDelete, pattern, handler)
}
func (a *HttpApp) Patch(pattern string, handler RouterHandler) {
	a.SetHandler(http.MethodPatch, pattern, handler)
}
func (a *HttpApp) Put(pattern string, handler RouterHandler) {
	a.SetHandler(http.MethodPut, pattern, handler)
}

func (a *HttpApp) Use(handler RouterHandler) {
	a.router.middlewares = append(a.router.middlewares, handler)
}
