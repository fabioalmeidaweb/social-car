package routes

import (
	"net/http"
	"social-car/src/controller"
)

var loginRoutes = []Route{
	{
		URI:     "/login",
		Method:  http.MethodPost,
		Handler: controller.Login,
		IsAuth:  false,
	},
}
