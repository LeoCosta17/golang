package repos

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func(repo Users) Search(identificador string)([]models.User, error){
	
	identificador = fmt.Sprintf("%%%s%%", identificador)

	results, err := repo.db.Query("SELECT id, nome, nick, email, dt_criado FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", 
		identificador, 
		identificador,
	)

	if err != nil{
		return nil, err
	}

	defer results.Close()

	var users []models.User

	for results.Next(){
		var user models.User

		if err = results.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreateDate); err != nil{
			return  nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func(repo Users) SearchByID(ID uint64)(models.User, error){

	result, err := repo.db.Query(
		"SELECT id, nome, nick, email, dt_criado FROM usuarios WHERE id = ?", 
		ID,
	)

	if err != nil{
		return models.User{}, err
	}	

	defer result.Close()

	var user models.User

	if result.Next(){
		if err = result.Scan(
			&user.ID, 
			&user.Name, 
			&user.Nick, 
			&user.Email, 
			&user.CreateDate,
		); err != nil{
			return models.User{}, err
		}
	}

	return user, nil
}