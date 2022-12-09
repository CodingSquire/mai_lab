package routers

import (
	"context"
	"net/http"
	"regexp"
	"users/controllers"
	"users/ctxkeys"
	m "users/middlewares"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routingTable   []route
	userController controllers.UserController
}

func NewRouter(u controllers.UserController) *Router {
	return &Router{
		userController: u,
	}
}

func (r *Router) setupRoutes() {
	r.routingTable = []route{
		{http.MethodGet, regexp.MustCompile(`^/users$`), r.userController.GetAllUsers},
		{http.MethodGet, regexp.MustCompile(`^/users/(?P<id>[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12})$`), r.userController.GetUserById},
		{http.MethodPost, regexp.MustCompile(`^/users$`), r.userController.CreateUser},
		{http.MethodPut, regexp.MustCompile(`^/users/(?P<id>[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12})$`), r.userController.UpdateUser},
		{http.MethodDelete, regexp.MustCompile(`^/users/(?P<id>[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12})$`), r.userController.DeleteUser},
	}
}

func (r *Router) serve(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routingTable {
		matches := route.regex.FindStringSubmatch(req.URL.Path)
		if route.method == req.Method && len(matches) > 0 {
			params := getParamsFromRoute(route, matches)
			ctx := context.WithValue(req.Context(), ctxkeys.ContextKeyParams, params)
			route.handler(w, req.WithContext(ctx))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func getParamsFromRoute(route route, matches []string) map[string]string {
	params := make(map[string]string)
	for i, name := range route.regex.SubexpNames() {
		if i != 0 && name != "" {
			params[name] = matches[i]
		}
	}
	return params
}

func (r *Router) Run(port string) {
	r.setupRoutes()
	http.Handle("/", m.Adapt(http.HandlerFunc(r.serve), m.LoggingMiddleware()))
	http.ListenAndServe(":"+port, nil)
}
