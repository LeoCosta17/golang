package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {

	utils.CarregarTemplates()
	router := router.Gerar()

	fmt.Println("Inicializando webapp na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
