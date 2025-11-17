package main

import (
	"crud/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rotas := routes.CarregaRotas()
	fs := http.FileServer(http.Dir("./static"))
	staticHandler := http.StripPrefix("/CRUD/static/", fs)

	rotas.Handle("GET /CRUD/static/", staticHandler)
	http.ListenAndServe(":5050", rotas)
}
