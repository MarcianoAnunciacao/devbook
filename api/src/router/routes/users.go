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
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/users/{id}",
		Method:                   http.MethodPut,
		Function:                 controllers.UpdateUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/users/{id}",
		Method:                   http.MethodDelete,
		Function:                 controllers.RemoveUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "users/{id}/follow",
		Method:                   http.MethodPost,
		Function:                 controllers.FollowAnUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "users/{id}/unfollow",
		Method:                   http.MethodPost,
		Function:                 controllers.FollowAnUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "users/{id}/followers",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchFollowers,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "users/{id}/following",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchUsersFollowedByAnUserID,
		IsAuthenticationRequired: true,
	},
}
