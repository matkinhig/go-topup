package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matkinhig/go-topup/middlewares"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := apiRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(
			middlewares.SetMiddleWareJSON(route.Handler)),
		).Methods(route.Method)
	}
	return r
}
