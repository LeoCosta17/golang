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

	router.HandleFunc("PUT /users/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarUsuario)))
	router.HandleFunc("DELETE /users/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.DeleteUsuario)))

	router.HandleFunc("POST /users/follow/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.SeguirUsuario)))
	router.HandleFunc("POST /users/unfollow/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.PararSeguirUsuario)))
	router.HandleFunc("GET /users/followers/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarSeguidores)))
	router.HandleFunc("GET /users/following/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.Seguindo)))

	router.HandleFunc("POST /users/update-password/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarSenha)))
	return router
}
