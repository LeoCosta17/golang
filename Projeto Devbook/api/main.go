package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("Rodando na porta: ", config.Port)
	fmt.Println("Dados acesso ao DB: ", config.ConnString)

	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
