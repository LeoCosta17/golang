package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var(
	// String para conexão com o banco
	ConnString = ""

	// Porta onde a API irá rodar
	Port = 0
)

// Carrega variáveis de ambiente
func Load(){
	var err error

	if err = godotenv.Load(); err != nil{
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil{
		Port = 9000
	}

	ConnString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}