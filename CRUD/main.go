package main

import (
	"crud/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rotas := routes.CarregaRotas()
	fs := http.FileServer(http.Dir("./static"))

	rotas.Handle("/CRUD/static/", http.StripPrefix("/CRUD/static/", fs))
	http.ListenAndServe(":5050", rotas)
}
