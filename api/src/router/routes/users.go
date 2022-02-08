package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:                      "/users",
		Method:                   http.MethodPost,
		Function:                 controllers.CreateUser,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/users",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchUsers,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/users/{id}",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchUser,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/users/{id}",
		Method:                   http.MethodPut,
		Function:                 controllers.UpdateUser,
		IsAuthenticationRequired: false,
	},
	{
		URI:                      "/users/{id}",
		Method:                   http.MethodDelete,
		Function:                 controllers.RemoveUser,
		IsAuthenticationRequired: false,
	},
}
