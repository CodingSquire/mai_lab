package http

import (
	"fmt"
	"net/http"
	"orders/table"
	"reflect"
	"regexp"
	"strings"
)

type HttpApp struct {
	Router Router
	State map[string]interface{}
}

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

func (h *HttpApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.RunPreMiddleware(w, r)
	h.Router.ServeHTTP(w, r, &h.State)
}

func (a *HttpApp) Run(port string) error {
	fmt.Println(a.MakeInfoTable(port))
	return http.ListenAndServe(":"+port, a)
}

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

func (a *HttpApp) Manage(state interface{}) {
	key := reflect.TypeOf(state).String()
	a.State[key] = state
}

func (a *HttpApp) SetHandler(method string, strPattern string, handler RouteHandler) {
	pattern := parsePattern(strPattern)

	newroutes := append(a.Router.routes, Route{
		OrigPath: strPattern,
		Method:   method,
		Pattern:  pattern,
		Handler:  handler,
	})

	a.Router.routes = newroutes
}

func (a *HttpApp) Get(pattern string, handler RouteHandler) {
	a.SetHandler(http.MethodGet, pattern, handler)
}
func (a *HttpApp) Post(pattern string, handler RouteHandler) {
	a.SetHandler(http.MethodPost, pattern, handler)
}
func (a *HttpApp) Delete(pattern string, handler RouteHandler) {
	a.SetHandler(http.MethodDelete, pattern, handler)
}
func (a *HttpApp) Patch(pattern string, handler RouteHandler) {
	a.SetHandler(http.MethodPatch, pattern, handler)
}
func (a *HttpApp) Put(pattern string, handler RouteHandler) {
	a.SetHandler(http.MethodPut, pattern, handler)
}

func (a *HttpApp) Use(handler MiddlwareRouterHandler) {
	a.Router.middlewares = append(a.Router.middlewares, handler)
}

func (a *HttpApp) Group(path string) *HttpGroup {
	return &HttpGroup {
		path: path,
		app: a,
	}
}

func (g *HttpGroup) Group(path string) *HttpGroup {
	return &HttpGroup {
		path: g.path + path,
		app: g.app,
	}
}

func (g *HttpGroup) Get(pattern string, handler RouteHandler) {
	g.app.SetHandler(http.MethodGet, g.path + pattern, handler)
}
func (g *HttpGroup) Post(pattern string, handler RouteHandler) {
	g.app.SetHandler(http.MethodPost, g.path + pattern, handler)
}
func (g *HttpGroup) Delete(pattern string, handler RouteHandler) {
	g.app.SetHandler(http.MethodDelete, g.path + pattern, handler)
}
func (g *HttpGroup) Patch(pattern string, handler RouteHandler) {
	g.app.SetHandler(http.MethodPatch, g.path + pattern, handler)
}
func (g *HttpGroup) Put(pattern string, handler RouteHandler) {
	g.app.SetHandler(http.MethodPut, g.path + pattern, handler)
}
