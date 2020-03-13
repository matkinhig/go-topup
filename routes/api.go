package routes

import (
	"net/http"

	"github.com/matkinhig/go-topup/controllers"
)

var apiRoutes = []Route{
	Route{
		Uri:     "/api/v1/deposit",
		Method:  http.MethodPost,
		Handler: controllers.DepositHandler,
	},
}
