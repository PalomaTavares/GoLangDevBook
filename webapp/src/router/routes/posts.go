package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{postID}/like",
		Method:                 http.MethodPut,
		Function:               controllers.LikePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{postID}/unlike",
		Method:                 http.MethodPut,
		Function:               controllers.UnlikePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post/{postID}/update",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUpdatePostPage,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{postID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{postID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		AuthenticationRequired: true,
	},
}
