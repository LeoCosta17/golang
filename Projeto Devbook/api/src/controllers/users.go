package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.ValidateUser(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewUsersRepos(db)

	user.ID, err = repo.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários..."))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

// Deleta um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}
