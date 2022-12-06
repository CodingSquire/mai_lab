package http

import (
	"context"
	"net/http"
	"regexp"
)

type RouteContext struct {
	params   map[string]string
	_request *http.Request
}

type RouterHandler func(w http.ResponseWriter, r *http.Request)

type Route struct {
	Method  string
	OrigPath string
	Pattern *regexp.Regexp
	Handler RouterHandler
}

type Router struct {
	middlewares []RouterHandler
	routes []Route
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

func (router *Router) RunPreMiddleware(w http.ResponseWriter, r *http.Request) {
	for _, middleware := range router.middlewares {
		middleware(w, r)
	}
} 

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		params, ok := getParams(*route.Pattern, r.URL.Path)

		if ok && r.Method == route.Method {
			ctx := context.WithValue(r.Context(), PARAMS, params)
			req := r.WithContext(ctx)
			route.Handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
