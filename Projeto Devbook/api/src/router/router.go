package router

import (
	"api/src/controllers"
	"net/http"
)

// Retorna um roteador com todas as rotas da api
func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /users", controllers.CriarUsuario)
	router.HandleFunc("GET /users/search/{user}", controllers.BuscarUsuariosPorIdentificador)
	router.HandleFunc("GET /users/search-id/{user_id}", controllers.BuscarUsuarioPorID)
	router.HandleFunc("PUT /users/update/{user_id}", controllers.AtualizarUsuario)
	router.HandleFunc("DELETE /users/delete/{user_id}", controllers.DeleteUsuario)

	return router
}
