package routes

import (
	"net/http"
	"social-car/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents all the routes of the API
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	IsAuth  bool
}

// Config configures the provided router with the necessary routes.
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoutes...)

	for _, route := range routes {

		if route.IsAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}

	}

	return r
}
