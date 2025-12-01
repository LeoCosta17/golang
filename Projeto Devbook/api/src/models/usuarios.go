package models

import (
	"errors"
	"strings"
	"time"
)

// Usuário da rede social
type Usuario struct {
	ID         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Senha      string    `json:"senha,omitempty"`
	DataCriado time.Time `json:"data_criado,omitempty"`
}

func (usuario *Usuario) ValidarUsuario(etapa string) error {
	if err := usuario.ValidarDados(etapa); err != nil {
		return err
	}

	usuario.Formatar()
	return nil
}

func (usuario *Usuario) ValidarDados(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco!")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco!")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco!")

	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco!")
	}

	return nil
}

func (usuario *Usuario) Formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
