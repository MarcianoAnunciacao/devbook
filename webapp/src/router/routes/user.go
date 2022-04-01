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
	{
		URI:                      "/users/{userId}/stop-following",
		Method:                   http.MethodPost,
		Function:                 controllers.StopFollowingUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/users/{userId}/following",
		Method:                   http.MethodPost,
		Function:                 controllers.FollowUser,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/profile",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadLoggedUserProfile,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/edit-user",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadUserProfileUpdate,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/edit-user",
		Method:                   http.MethodPut,
		Function:                 controllers.EditUserProfile,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/update-password",
		Method:                   http.MethodGet,
		Function:                 controllers.LoadUpdatePasswordPage,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/update-password",
		Method:                   http.MethodPost,
		Function:                 controllers.UpdatePassword,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/delete-user",
		Method:                   http.MethodDelete,
		Function:                 controllers.DeleteUser,
		IsAuthenticationRequired: true,
	},
}
