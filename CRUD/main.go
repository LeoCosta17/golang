package main

import (
	"crud/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.ListenAndServe(":5050", routes.CarregaRotas())
}
