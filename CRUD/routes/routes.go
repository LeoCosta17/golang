package routes

import (
	"crud/controllers"
	"net/http"
)

func CarregaRotas() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /home", controllers.HomeHandler)

	return router
}
