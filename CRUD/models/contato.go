package models

import (
	"crud/database"
	"fmt"
)

type Contato struct {
	Nome     string
	Email    string
	Telefone string
}

func (Contato) CriarContato(contato *Contato) error {

	db, err := database.ConnDatabase()
	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuarios(nome, email, telefone) VALUES ($1, $2, $3)")

	if err != nil {
		return err
	}

	defer statement.Close()

	resultado, err := statement.Exec(contato.Nome, contato.Email, contato.Telefone)
	if err != nil {
		return err
	}

	_, err = resultado.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("Contato criado:", contato.Nome, contato.Email, contato.Telefone)
	return nil
}

func (Contato) BuscarContatos() ([]Contato, error) {
	db, err := database.ConnDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return nil, nil
}

func (Contato) BuscarContato() (*Contato, error) {
	return nil, nil
}
