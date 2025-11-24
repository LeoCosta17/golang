package repos

import (
	"api/src/models"
	"database/sql"
)

type Users struct{
	db *sql.DB
}

func NewUsersRepos(db *sql.DB) *Users{
	return &Users{db}
}

func (repo Users) Create(user models.User)(uint64, error){

	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha)VALUES (?, ?, ?, ?)",
	)

	if err != nil{
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.PassWord)

	if err != nil{
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func(repo Users) Search()([]models.User, error){
	results, err := repo.db.Query("SELECT * FROM usuarios")

	if err != nil{
		return []models.User{} ,err
	}

	defer results.Close()

	var users []models.User

	for results.Next(){
		var user models.User

		if err := results.Scan() ;err != nil{
			return []models.User{}, err
		}

		users = append(users, user)

	}

	return users, nil
}