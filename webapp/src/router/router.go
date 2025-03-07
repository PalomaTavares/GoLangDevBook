package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// router with all configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
