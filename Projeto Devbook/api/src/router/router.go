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
	router.HandleFunc("POST /usuarios", middlewares.Logger(controllers.CriarUsuario))

	router.HandleFunc("GET /usuarios/buscar/{user}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarUsuariosPorIdentificador)))
	router.HandleFunc("GET /usuarios/buscar-id/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarUsuarioPorID)))

	router.HandleFunc("PUT /usuarios/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarUsuario)))
	router.HandleFunc("DELETE /usuarios/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.DeleteUsuario)))

	router.HandleFunc("POST /usuarios/seguir/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.SeguirUsuario)))
	router.HandleFunc("POST /usuarios/deixar-seguir/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.PararSeguirUsuario)))
	router.HandleFunc("GET /usuarios/seguidores/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarSeguidores)))
	router.HandleFunc("GET /usuarios/seguindo/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.Seguindo)))

	router.HandleFunc("POST /usuarios/atualizar-senha/{user_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarSenha)))

	router.HandleFunc("POST /publicacoes", middlewares.Logger(middlewares.Autenticar(controllers.CriarPublicacao)))
	router.HandleFunc("GET /publicacoes", middlewares.Logger(middlewares.Autenticar(controllers.BuscarPublicacoes)))
	router.HandleFunc("GET /publicacoes/buscar/{publicacao_id}", middlewares.Logger(middlewares.Autenticar(controllers.BuscarPublicacao)))
	router.HandleFunc("PUT /publicacoes/atualizar/{publicacao_id}", middlewares.Logger(middlewares.Autenticar(controllers.AtualizarPublicacao)))
	router.HandleFunc("DELETE /publicacoes/deletar/{publicacao_id}", middlewares.Logger(middlewares.Autenticar(controllers.DeletarPublicacao)))

	return router
}
