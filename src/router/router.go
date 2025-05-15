package router

import (
	"main/src/router/routes"

	"github.com/gorilla/mux"
)

func RunServer() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
