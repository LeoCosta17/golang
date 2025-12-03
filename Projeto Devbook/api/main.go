package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
	func init() {
		chave := make([]byte, 64)
		if _, err := rand.Read(chave); err != nil {
			log.Fatal(err)
		}
		fmt.Println(chave)

		string64 := base64.StdEncoding.EncodeToString(chave)
		fmt.Println(string64)
	}
*/
func main() {
	config.Load()
	fmt.Println("Rodando na porta: ", config.Port)
	fmt.Println("Dados acesso ao DB: ", config.ConnString)
	//fmt.Println(config.SecretKey)
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
