package router

import (
	"net/http"
	"web/src/controllers"
)

// Retorna um roteador com todas as rotas configuradas
func Gerar() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controllers.CarregarTelaLogin)
	router.HandleFunc("GET /login", controllers.CarregarTelaLogin)
	router.HandleFunc("POST /login", controllers.FazerLogin)

	router.HandleFunc("GET /criar-usuario", controllers.CarregarPaginaCadastroUsuario)
	router.HandleFunc("POST /usuarios", controllers.CriarUsuario)

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer))
	return router
}
