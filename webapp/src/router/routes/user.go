package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                      "/create-user",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadCreateUserPage,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/users",
		Method:                   http.MethodPost,
		Function:                 controllers.CreateUser,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/search-users",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadUsersPage,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/users/{userId}",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadUserProfile,
		IsAuthenticationRequired: true,
	},
}
