package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	fmt.Println("Inicializando webapp...")

	router := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", router))
}
