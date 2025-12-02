package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if err := usuario.Formatar(etapa); err != nil {
		return err
	}
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

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("Email inserido é inválido!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco!")
	}

	return nil
}

func (usuario *Usuario) Formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senha_criptografada, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senha_criptografada)
	}

	return nil
}
