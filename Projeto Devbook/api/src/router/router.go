package router

import (
	"api/src/controllers"
	"net/http"
)

// Retorna um roteador com todas as rotas da api
func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /users", controllers.CreateUser)
	router.HandleFunc("GET /users", controllers.SearchUsers)
	router.HandleFunc("GET /users/{user_id}", controllers.SearchUser)
	router.HandleFunc("PUT /users/{user_id}", controllers.UpdateUser)
	router.HandleFunc("DELETE /users/{user_id}", controllers.DeleteUser)


	return router
}
