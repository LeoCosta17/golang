package routes

import (
	"crud/controllers"
	"net/http"
)

func CarregaRotas() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controllers.HomeHandler)
	router.HandleFunc("POST /", controllers.CriarContatoHandler)
	//router.HandleFunc("PUT /", nil)
	router.HandleFunc("DELETE /{id}", controllers.ExcluirContatoHandler)

	return router
}
