package http

import (
	"fmt"
	"net/http"
	"orders/table"
	"reflect"
)

// Routing base interface for some sub routing
type HttpRouter interface {
	Get(pattern string, handler RouteHandler)
	Post(pattern string, handler RouteHandler)
	Delete(pattern string, handler RouteHandler)
	Patch(pattern string, handler RouteHandler)
	Put(pattern string, handler RouteHandler)
	Use(handlers ...RouteHandler)
	Group(pattern string) HttpRouter
}

// Main app structure with inner map of global-like objects
type HttpApp struct {
	Router Router
	State map[string]interface{}
}

// Group for inner nesting
type HttpGroup struct {
	path string
	app *HttpApp
}

func NewApp() *HttpApp {
	return &HttpApp{
		Router: Router{},
		State: make(map[string]interface{}),
	}
}

// Entrypoint of server
func (h *HttpApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r, &h.State)
}

// Running server on some port
// TODO: add host setting option
func (a *HttpApp) Run(port string) error {
	fmt.Println(a.MakeInfoTable(port))
	return http.ListenAndServe(":"+port, a)
}

// Just printing fancy table of routes on router
func (a *HttpApp) MakeInfoTable(port string) string {
	var table table.TableBuilder

	for _, route := range a.Router.routes {
		table.AppendRoute(route.Method, route.OrigPath)
	}

	table.AppendLine("Runtime:")
	table.AppendLine("localhost:"+port)

	table.PrependLine("Routes:")

	return table.String()
}

// Addig global-like object with somewhat reflection
func (a *HttpApp) Manage(state interface{}) {
	key := reflect.TypeOf(state).String()
	a.State[key] = state
}

// Main route methods on main app

func (a *HttpApp) Get(pattern string, handler RouteHandler) {
	a.Router.SetHandler(http.MethodGet, pattern, handler)
}
func (a *HttpApp) Post(pattern string, handler RouteHandler) {
	a.Router.SetHandler(http.MethodPost, pattern, handler)
}
func (a *HttpApp) Delete(pattern string, handler RouteHandler) {
	a.Router.SetHandler(http.MethodDelete, pattern, handler)
}
func (a *HttpApp) Patch(pattern string, handler RouteHandler) {
	a.Router.SetHandler(http.MethodPatch, pattern, handler)
}
func (a *HttpApp) Put(pattern string, handler RouteHandler) {
	a.Router.SetHandler(http.MethodPut, pattern, handler)
}
func (a *HttpApp) Use(handlers ...RouteHandler) {
	for _, handler := range handlers {
		a.Router.SetHandler("*", "/",  handler)
	}
}


// Gouping for subnesting
func (a *HttpApp) Group(path string) HttpRouter {
	return &HttpGroup {
		path: path,
		app: a,
	}
}

// Same methods as fot HttpApp with route concatination

func (g *HttpGroup) Group(path string) HttpRouter {
	return &HttpGroup {
		path: g.path + path,
		app: g.app,
	}
}

func (g *HttpGroup) Use(handlers ...RouteHandler) {
	for _, handler := range handlers {
		g.app.Router.SetHandler("*", g.path,  handler)
	}
}
func (g *HttpGroup) Get(pattern string, handler RouteHandler) {
	g.app.Get(g.path + pattern, handler)
}
func (g *HttpGroup) Post(pattern string, handler RouteHandler) {
	g.app.Post(g.path + pattern, handler)
}
func (g *HttpGroup) Delete(pattern string, handler RouteHandler) {
	g.app.Delete(g.path + pattern, handler)
}
func (g *HttpGroup) Patch(pattern string, handler RouteHandler) {
	g.app.Patch(g.path + pattern, handler)
}
func (g *HttpGroup) Put(pattern string, handler RouteHandler) {
	g.app.Put(g.path + pattern, handler)
}
