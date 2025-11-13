package routes

import (
	"crud/controllers"
	"net/http"
)

func CarregaRotas() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /home", controllers.HomeHandler)
	router.HandleFunc("POST /home", controllers.CriarContatoHandler)
	//router.HandleFunc("PUT /home", nil)
	//router.HandleFunc("DELETE /home", nil)

	return router
}
