package router

import (
	"social-car/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a new instance of the mux.Router.
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
