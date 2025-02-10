package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequireAuthetication: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.GetAllUsers,
		RequireAuthetication: true,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodGet,
		Function:             controllers.GetUser,
		RequireAuthetication: true,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodPut,
		Function:             controllers.UpdateUser,
		RequireAuthetication: true,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequireAuthetication: true,
	},
}
