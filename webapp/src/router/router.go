package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
