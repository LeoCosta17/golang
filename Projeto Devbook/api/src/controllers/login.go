package controllers

import (
	"api/src/autenticacao"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"net/http"
)

// Responsável pela autenticação de usuário na aplicação
func Login(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)

	usuarioBuscado, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(usuarioBuscado.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := autenticacao.CriarToken(usuarioBuscado.ID)

	w.Write([]byte(token))
}
