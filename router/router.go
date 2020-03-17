package router

import (
	"github.com/gorilla/mux"
	"github.com/matkinhig/go-topup/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
