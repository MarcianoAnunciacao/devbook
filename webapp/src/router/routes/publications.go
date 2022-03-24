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
		Method:                   http.MethodGet,
		Function:                 controllers.LoadPublicationEditionPage,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{publicationId}",
		Method:                   http.MethodPut,
		Function:                 controllers.UpdatePublication,
		IsAuthenticationRequired: true,
	},
	{
		URI:                      "/publications/{publicationId}",
		Method:                   http.MethodDelete,
		Function:                 controllers.DeletePublication,
		IsAuthenticationRequired: true,
	},
}
