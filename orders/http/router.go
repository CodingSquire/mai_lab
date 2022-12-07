package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
)

type RouteContext struct {
	w      http.ResponseWriter
	params map[string]string
	state  *map[string]interface{}
	R      *http.Request
}

type ResponseWriter http.ResponseWriter

type RouteHandler func(r *RouteContext)
type MiddlwareRouterHandler func(w ResponseWriter, r *http.Request)

type Route struct {
	Method   string
	OrigPath string
	Pattern  *regexp.Regexp
	Handler  RouteHandler
}

type Router struct {
	middlewares []MiddlwareRouterHandler
	routes      []Route
}

func (r *RouteContext) Params(key string) string {
	return r.params[key]
}

func (r *RouteContext) State(toGet interface{}) {
	targetType := reflect.TypeOf(toGet)

	d := (*r.state)[targetType.String()]

	//Actually set found struct
	reflect.ValueOf(toGet).Elem().Set(reflect.ValueOf(d).Elem())
}

func (r *RouteContext) Body() io.ReadCloser {
	return r.R.Body
}

func (r *RouteContext) SendError(error error) {
	http.Error(r.w, error.Error(), http.StatusBadRequest)
}

func (r *RouteContext) SendString(message string) {
	fmt.Fprint(r.w, message)
}

func (r *RouteContext) SendJSON(obj interface{}) {
	r.R.Header.Set("Content-Type", "application/json")
	json.NewEncoder(r.w).Encode(obj)
}

func getParams(regEx regexp.Regexp, url string) (map[string]string, bool) {
	match := regEx.FindStringSubmatch(url)
	paramsMap := make(map[string]string)

	if match == nil {
		return paramsMap, false
	}

	for i, name := range regEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap, true
}

func (router *Router) RunPreMiddleware(w http.ResponseWriter, r *http.Request) {
	for _, middleware := range router.middlewares {
		middleware(w, r)
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request, state *map[string]interface{}) {
	for _, route := range router.routes {
		params, ok := getParams(*route.Pattern, r.URL.Path)

		if ok && r.Method == route.Method {
			route.Handler(&RouteContext{
				w:      w,
				params: params,
				state:  state,
				R:      r,
			})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
