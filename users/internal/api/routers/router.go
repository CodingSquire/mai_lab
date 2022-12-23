// Package routers contains the main router for the application.
package routers

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"users/internal/api/common"
	"users/internal/api/controllers"
	"users/internal/api/middlewares"
)

// route is a single route in the routing table.
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

// Router is the main router for the application.
type Router struct {
	routingTable    []route
	userController  controllers.UserController
	orderController controllers.OrderController
}

// NewRouter creates a new router with the given user controller.
func NewRouter(u controllers.UserController, o controllers.OrderController) *Router {
	return &Router{
		userController:  u,
		orderController: o,
	}
}

// idGroup is a regex group for matching UUIDs.
const idGroup = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"

// setupRoutes sets up the routing table.
func (r *Router) setupRoutes() {
	r.routingTable = []route{
		{http.MethodGet, regexp.MustCompile(`^/users$`), r.userController.GetAllUsers},
		{http.MethodGet, regexp.MustCompile(`^/users/(?P<id>` + idGroup + `)$`), r.userController.GetUserById},
		{http.MethodPost, regexp.MustCompile(`^/users$`), r.userController.CreateUser},
		{http.MethodPut, regexp.MustCompile(`^/users/(?P<id>` + idGroup + `)$`), r.userController.UpdateUser},
		{http.MethodDelete, regexp.MustCompile(`^/users/(?P<id>` + idGroup + `)$`), r.userController.DeleteUser},
		{http.MethodPost, regexp.MustCompile(`^/users/(?P<id>` + idGroup + `)/orders$`), r.orderController.CreateOrderByUserId},
		{http.MethodGet, regexp.MustCompile(`^/users/(?P<id>` + idGroup + `)/orders$`), r.orderController.GetAllOrdersByUserId},
		{http.MethodGet, regexp.MustCompile(`^/orders/(?P<orderId>` + idGroup + `)$`), r.orderController.GetOrderById},
	}
}

// serve is the main handler for the router.
func (r *Router) serve(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routingTable {
		matches := route.regex.FindStringSubmatch(req.URL.Path)
		if route.method == req.Method && len(matches) > 0 {
			params := getParamsFromRoute(route, matches)
			ctx := context.WithValue(req.Context(), common.ContextKeyParams, params)
			route.handler(w, req.WithContext(ctx))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

// getParamsFromRoute returns a map of the route parameters.
func getParamsFromRoute(route route, matches []string) map[string]string {
	params := make(map[string]string)
	for i, name := range route.regex.SubexpNames() {
		if i != 0 && name != "" {
			params[name] = matches[i]
		}
	}
	return params
}

// Run starts the router.
func (r *Router) Run(port string) {
	r.setupRoutes()
	http.Handle("/", middlewares.Adapt(http.HandlerFunc(r.serve), middlewares.LoggingMiddleware()))
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
