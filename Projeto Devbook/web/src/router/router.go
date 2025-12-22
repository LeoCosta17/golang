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

	return router
}
