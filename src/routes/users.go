package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CrateUser,
		RequireAuthetication: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.GetAllUsers,
		RequireAuthetication: false,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.GetUser,
		RequireAuthetication: false,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdateUser,
		RequireAuthetication: false,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequireAuthetication: false,
	},
}
