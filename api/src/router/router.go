package router

import "github.com/gorilla/mux"

// It should returns a configured router
func Gerar() *mux.Router {
	return mux.NewRouter()
}
