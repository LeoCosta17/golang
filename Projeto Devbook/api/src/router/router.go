package router

import (
	"api/src/controllers"
	"net/http"
)

// Retorna um roteador com todas as rotas da api
func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /users", controllers.CreateUser)
	router.HandleFunc("GET /users/search/{user}", controllers.SearchUsersByIdentifier)
	router.HandleFunc("GET /users/search-id/{user_id}", controllers.SearchUserByID)
	router.HandleFunc("PUT /users/{user_id}", controllers.UpdateUser)
	router.HandleFunc("DELETE /users/{user_id}", controllers.DeleteUser)


	return router
}
