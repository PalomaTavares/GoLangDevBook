package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                  "/posts",
		Method:               http.MethodPost,
		Function:             controllers.CreatePost,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts",
		Method:               http.MethodGet,
		Function:             controllers.GetAllPosts,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts/{postID}",
		Method:               http.MethodGet,
		Function:             controllers.GetPost,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts/{postID}",
		Method:               http.MethodPut,
		Function:             controllers.UpdatePost,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts/{postID}",
		Method:               http.MethodDelete,
		Function:             controllers.DeletePost,
		RequireAuthetication: true,
	},
	{
		URI:                  "/users/{userID}/posts",
		Method:               http.MethodGet,
		Function:             controllers.GetPostsByUserID,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts/{postID}/like",
		Method:               http.MethodPut,
		Function:             controllers.Like,
		RequireAuthetication: true,
	},
	{
		URI:                  "/posts/{postID}/unlike",
		Method:               http.MethodPut,
		Function:             controllers.Unlike,
		RequireAuthetication: true,
	},
}
