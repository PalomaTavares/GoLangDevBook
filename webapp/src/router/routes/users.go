package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadNewAccountPage,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/search-users",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersPage,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersProfilePage,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/profile",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoggedUserProfile,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/edit-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadEditUserProfile,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/edit-user",
		Method:                 http.MethodPut,
		Function:               controllers.EditUserProfile,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUpdatePassword,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/delete-user",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: true,
	},
}
