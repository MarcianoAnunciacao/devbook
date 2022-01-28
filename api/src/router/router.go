package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// It should returns a configured router
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
