package routes

import "net/http"

var publicationRoutes = []Route{
	{
		URI:                      "/publications",
		Method:                   http.MethodPost,
		Function:                 controllers.CreatePublication,
		IsAuthenticationRequired: true,
	},
}
