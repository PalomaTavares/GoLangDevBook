package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homePageRoute = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.LoadHomePage,
	AuthenticationRequired: true,
}
