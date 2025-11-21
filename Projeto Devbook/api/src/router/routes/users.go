package routes

import "net/http"

var usersRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
}
