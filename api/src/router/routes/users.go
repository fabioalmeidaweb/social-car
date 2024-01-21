package routes

import (
	"net/http"
	"social-car/src/controller"
)

var usersRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controller.CreateUser,
		IsAuth:  false,
	},
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controller.GetUsers,
		IsAuth:  true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodGet,
		Handler: controller.GetUser,
		IsAuth:  true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodPut,
		Handler: controller.UpdateUser,
		IsAuth:  true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodDelete,
		Handler: controller.DeleteUser,
		IsAuth:  true,
	},
	{
		URI:     "/users/{userId}/follow",
		Method:  http.MethodPost,
		Handler: controller.FollowUser,
		IsAuth:  true,
	}, {
		URI:     "/users/{userId}/unfollow",
		Method:  http.MethodPost,
		Handler: controller.UnFollowUser,
		IsAuth:  true,
	},
}
