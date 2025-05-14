package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

var healthRoute = Route{
	Uri:    "/",
	Method: http.MethodGet,
	Function: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	},
	AuthenticationRequired: false,
}

func Config(r *mux.Router) *mux.Router {
	var routes []Route

	routes = append(routes, healthRoute)
	routes = append(routes, UserRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
