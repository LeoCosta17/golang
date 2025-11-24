package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em JSON para uma request
func JSON(w http.ResponseWriter, statusCode int, data interface{}){
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil{
		log.Fatal(err)
	}
}

// Retorna um erro em formato JSON
func Error(w http.ResponseWriter, statusCode int, err error){
	JSON(w, statusCode, struct{
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}