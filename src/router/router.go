package router

import (
	"api/src/routes"

	"github.com/gorilla/mux"
)

// retuns a router with all routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
