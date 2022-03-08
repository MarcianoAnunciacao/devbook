package router

import "github.com/gorilla/mux"

func GenerateRoutes() *mux.Router {
	return mux.NewRouter()
}
