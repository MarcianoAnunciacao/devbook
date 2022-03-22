package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var publicationRoutes = []Route{
	{
		URI:                      "/publications",
		Method:                   http.MethodPost,
		Function:                 controllers.CreatePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{publicationId}/like",
		Method:                   http.MethodPost,
		Function:                 controllers.LikePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{publicationId}/dislike",
		Method:                   http.MethodPost,
		Function:                 controllers.DislikePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{publicationId}/edit",
		Method:                   http.MethodPost,
		Function:                 controllers.DislikePublication,
		IsAuthenticationRequired: true,
	},
}
