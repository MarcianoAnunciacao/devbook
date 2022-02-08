package routes

import (
	"api/src/midlewares"
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
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.IsAuthenticationRequired {
			r.HandleFunc(
				route.URI,
				midlewares.Logger(midlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, midlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
