package router

import "net/http"

// Retorna um roteador com todas as rotas da api
func CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	return router
}
