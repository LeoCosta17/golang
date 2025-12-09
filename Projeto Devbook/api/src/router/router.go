package router

import (
	"api/src/controllers"
	"api/src/middlewares"
	"net/http"
)

// Retorna um roteador com todas as rotas da api
func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /login", controllers.Login)
	router.HandleFunc("POST /users", middlewares.Logger(controllers.CriarUsuario))
	router.HandleFunc("GET /users/search/{user}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarUsuariosPorIdentificador)))
	router.HandleFunc("GET /users/search-id/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarUsuarioPorID)))
	router.HandleFunc("PUT /users/update/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarUsuario)))
	router.HandleFunc("DELETE /users/delete/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.DeleteUsuario)))

	return router
}
