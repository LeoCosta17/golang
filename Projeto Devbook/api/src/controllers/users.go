package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"encoding/json"
	"net/http"
	"strconv"
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

func SearchUsersByIdentifier(w http.ResponseWriter, r *http.Request) {

	identificador := r.PathValue("user")

	db, err := database.DBConn()
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repos.NewUsersRepos(db)

	users, err := repo.Search(identificador)
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)	
}

func SearchUserByID(w http.ResponseWriter, r *http.Request) {
	
	ID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)

	if err != nil{
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DBConn()

	repo := repos.NewUsersRepos(db)

	user, err := repo.SearchByID(ID)

	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

// Deleta um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}
