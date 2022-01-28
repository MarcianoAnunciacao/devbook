package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represents all APIs routes
type Routes struct {
	URI                      string
	Method                   string
	Function                 func(http.ResponseWriter, *http.Request)
	IsAuthenticationRequired bool
}

//Set all routes inside a router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
