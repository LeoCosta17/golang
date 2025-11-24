package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	
	if r.Method != http.MethodPost{
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err !=nil{
		log.Fatal(err)
	}

	db, err := database.DBConn()
	if err != nil{
		log.Fatal(err)
	}

	repo := repos.NewUsersRepos(db)

	userID, err := repo.Create(user)

	if err != nil{
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Criado usuário nº: %d", userID)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuários..."))
}


func SearchUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuário..."))
}


func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando usuário..."))
}

// Deleta um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando usuário..."))
}