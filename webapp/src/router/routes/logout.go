package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRote = Route{
	URI:                      "/logout",
	Method:                   http.MethodGet,
	Function:                 controllers.Logout,
	IsAuthenticationRequired: true,
}
