package models

import (
	"errors"
	"strings"
	"time"
)

// Usuário da rede social
type User struct{
	ID uint64 `json:"id,omitempty"`
	Name string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	PassWord string `json:"passWord,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

func (user *User) ValidateUser() error{
	if err := user.ValidateData(); err != nil{
		return err
	}

	user.Format()
	return nil
}

func (user *User) ValidateData() error{
	if user.Name == ""{
		return errors.New("O nome é obrigatório e não pode estar em branco!")
	}
	if user.Nick == ""{
		return errors.New("O nick é obrigatório e não pode estar em branco!")
	}
	if user.Email == ""{
		return errors.New("O email é obrigatório e não pode estar em branco!")
	}
	if user.PassWord == ""{
		return errors.New("A senha é obrigatória e não pode estar em branco!")
	}

	return nil
}

func (user *User) Format(){
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}