package models

import "time"

// Usuário da rede social
type User struct{
	ID uint32 `json:"id,omitempty"`
	Name string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	PassWord string `json:"passWord, omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

