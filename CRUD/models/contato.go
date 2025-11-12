package models

import "fmt"

type Contato struct {
	Nome     string
	Email    string
	Telefone string
}

func (Contato) CriarContato(contato *Contato) error {

	fmt.Println("Contato criado:", contato.Nome, contato.Email, contato.Telefone)
	return nil
}
