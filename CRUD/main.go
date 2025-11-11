package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /home", Home)

	http.ListenAndServe(":5050", router)
}
