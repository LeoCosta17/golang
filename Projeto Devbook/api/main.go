package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.CreateRouter()

	fmt.Println("Servidor rodando porta 6000!")
	log.Fatal(http.ListenAndServe(":6000", router))
}
