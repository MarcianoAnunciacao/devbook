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
}
