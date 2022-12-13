package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

// RouteContext - context of route with params, and so on
type RouteContext struct {
	W      http.ResponseWriter
	R      *http.Request
	Router *Router
	// Parsed params
	params map[string]string
	state  *map[string]interface{}
	index  int
}

type ResponseWriter http.ResponseWriter

type RouteHandler func(r *RouteContext)

// Route info patterns/handlers/etc
type Route struct {
	Method   string
	OrigPath string
	Pattern  *regexp.Regexp
	Handler  RouteHandler
}

type Router struct {
	routes []*Route
}


// Next - walking on routes till error or handler
func (r *RouteContext) Next() error {
	// TODO: can add may handlers on route
	// and we can walk them here
	return r.Router.next(r)
}

// Params returns param on route
func (r *RouteContext) Params(key string) string {
	return r.params[key]
}

// State is getting global-like objects by somewhat reflection
func (r *RouteContext) State(toGet interface{}) {
	targetType := reflect.TypeOf(toGet)

	d := (*r.state)[targetType.String()]

	//Actually set found struct
	reflect.ValueOf(toGet).Elem().Set(reflect.ValueOf(d).Elem())
}

// Body getting body of request
func (r *RouteContext) Body() io.ReadCloser {
	return r.R.Body
}

// DecodeJSON is a wrapper of json decoder
func (r *RouteContext) DecodeJSON(obj any) error {
	return json.NewDecoder(r.Body()).Decode(obj)
}


// SendError making response with error
func (r *RouteContext) SendError(error error) {
	http.Error(r.W, error.Error(), http.StatusBadRequest)
}

// SendString setting string message on response
func (r *RouteContext) SendString(message string) {
	fmt.Fprint(r.W, message)
}

// SendJSON encode obj as json on response
func (r *RouteContext) SendJSON(obj interface{}) {
	r.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r.W).Encode(obj)
}

// Parsing params by regex
func (r *Route) getParams(url string) (map[string]string, bool) {
	match := r.Pattern.FindStringSubmatch(url)
	paramsMap := make(map[string]string)

	if match == nil {
		return paramsMap, false
	}

	for i, name := range r.Pattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap, true
}

// Simple glob check
func (r *Route) methodIsValidAgainst(method string) bool {
	return r.Method == "*" || method == r.Method
}

// ServeHTTP is a main http handler
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request, state *map[string]interface{}) {
	ctx := &RouteContext{
		W:      w,
		R:      r,
		Router: router,
		state:  state,
		index: -1,
	}
	err := router.next(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// Main walk function where we find next sutable route to handle
func (router *Router) next(ctx *RouteContext) (err error) {
	for ctx.index + 1 < len(router.routes) {
		// getting new route
		ctx.index++
		route := router.routes[ctx.index]

		// parsing params
		params, ok := route.getParams(ctx.R.URL.Path)

		if ok && route.methodIsValidAgainst(ctx.R.Method) {
			// setting parsed params
			ctx.params = params
			route.Handler(ctx)
			return
		}
	}
	err = fmt.Errorf("No route found")
	// XXX: shoud we write error to ctx.W?
	return
}

// Parsing predefined route patterns
// Example: '/route/:id', etc...
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

// SetHandler is setting new handler on route pattern
func (r *Router) SetHandler(method string, pathRaw string, handler RouteHandler) {
	// Cannot have an empty path
	if pathRaw == "" {
		pathRaw = "/"
	}
	// Path always start with a '/'
	if pathRaw[0] != '/' {
		pathRaw = "/" + pathRaw
	}

	// fallback on glob
	if method == "" {
		method = "*"
	}

	// TODO: handle globs in routes && case-insensetive

	pattern := parsePattern(pathRaw)

	newroutes := append(r.routes, &Route{
		OrigPath: pathRaw,
		Method:   method,
		Pattern:  pattern,
		Handler:  handler,
	})

	r.routes = newroutes
}
