package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Routes{
	{
		URI:                      "/publications",
		Method:                   http.MethodPost,
		Function:                 controllers.CreatePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchPublications,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{id}",
		Method:                   http.MethodGet,
		Function:                 controllers.SearchPublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{id}",
		Method:                   http.MethodPut,
		Function:                 controllers.UpdatePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{id}",
		Method:                   http.MethodDelete,
		Function:                 controllers.DeletePublication,
		IsAuthenticationRequired: true,
	},
}
