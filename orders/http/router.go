package http

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

type RouteContext struct {
	params   map[string]string
	_request *http.Request
}

type RouterHandler func(w http.ResponseWriter, r *http.Request)

type innerRouter struct {
	method  string
	origPath string
	pattern *regexp.Regexp
	handler RouterHandler
}

type Router struct {
	middlewares []RouterHandler
	routes []innerRouter
}

const PARAMS = "params"

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

const COLOR_DEFAULT = "\033[39m"
const COLOR_RED = "\033[91m"
const COLOR_GREEN = "\033[92m"
const COLOR_BLUE = "\033[94m"
const COLOR_WHITE = "\033[97m"

func getColorByMethod(method string) string {
	switch method {
	case http.MethodGet:
		return string(COLOR_GREEN)
	case http.MethodDelete:
		return string(COLOR_RED)
	case http.MethodPost:
		return string(COLOR_BLUE)
	default:
		return string(COLOR_WHITE)
	}
}

func (r *Router) PrettyPrint() {
	fmt.Println("Setted routes:")
	for _, route := range r.routes {
		color := getColorByMethod(route.method)
		fmt.Printf(
			"%s  %-6s%s\t%s\n",
			color,
			route.method,
			COLOR_DEFAULT,
			route.origPath,
		)
	}
}

func (router *Router) RunPreMiddleware(w http.ResponseWriter, r *http.Request) {
	for _, middleware := range router.middlewares {
		middleware(w, r)
	}
} 

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		params, ok := getParams(*route.pattern, r.URL.Path)

		if ok && r.Method == route.method {
			ctx := context.WithValue(r.Context(), PARAMS, params)
			req := r.WithContext(ctx)
			route.handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
