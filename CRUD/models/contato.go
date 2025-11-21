package models

import (
	"crud/database"
	"fmt"
)

type Contato struct {
	ID       int
	Nome     string
	Email    string
	Telefone string
}

func (Contato) CriarContato(contato *Contato) error { // Cadastra um contato no banco de dados

	db, err := database.ConnDatabase()
	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare("INSERT INTO contatos(nome, email, telefone) VALUES ($1, $2, $3)")

	if err != nil {
		return err
	}

	defer statement.Close()

	resultados, err := statement.Exec(contato.Nome, contato.Email, contato.Telefone)
	if err != nil {
		return err
	}

	_, err = resultados.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("Contato criado:", contato.Nome, contato.Email, contato.Telefone)
	return nil
}

func BuscarContatos() ([]Contato, error) { // Busca no banco todos os contatos cadastrados e retorna um slice contendo todos eles
	db, err := database.ConnDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	resultados, err := db.Query("SELECT * FROM contatos")
	if err != nil {
		return nil, err
	}

	defer resultados.Close()

	var contatos []Contato

	for resultados.Next() {
		var contato Contato

		if err := resultados.Scan(&contato.ID, &contato.Nome, &contato.Email, &contato.Telefone); err != nil {
			return nil, err
		}

		contatos = append(contatos, contato)
	}

	return contatos, nil
}

func BuscarContato() (*Contato, error) { // Busca por um contato específico no banco de dados e retorna o mesmo
	return nil, nil
}

func (Contato) EditarContato(contato *Contato) error { // Edita um contato específico no banco de dados
	db, err := database.ConnDatabase()
	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare("UPDATE contatos SET nome=$1, email=$2, telefone=$3 WHERE id=$4")
	if err != nil {
		return err
	}

	defer statement.Close()

	resultado, err := statement.Exec(contato.Nome, contato.Email, contato.Telefone, contato.ID)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func ExcluirContato(ID_contato int) error { // Exclui um contato específico no banco de dados

	db, err := database.ConnDatabase()
	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare("DELETE FROM contatos WHERE id=$1")
	if err != nil {
		return err
	}
	defer statement.Close()

	resultado, err := statement.Exec(ID_contato)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
