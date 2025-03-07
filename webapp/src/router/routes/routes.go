package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// web aplication routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, userRoutes...)
	routes = append(routes, homePageRoute)
	routes = append(routes, postRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {
		if route.AuthenticationRequired {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileServer))
	return router
}
