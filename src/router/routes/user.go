package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
