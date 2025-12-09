package controllers

import (
	"api/src/autenticacao"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := usuario.ValidarUsuario("cadastro"); err != nil {
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

	usuario.ID, err = repositorio.Criar(usuario)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario.ID)
}

// Busca um usuário por um identificador dele: Nome ou Nick
func BuscarUsuariosPorIdentificador(w http.ResponseWriter, r *http.Request) {

	identificador := r.PathValue("user")

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)

	usuarios, err := repositorio.Buscar(identificador)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// Busca um usuário por seu ID
func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {

	ID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)

	if err != nil {
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

	usuario, err := repositorio.BuscarPorID(ID)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

// Atualiza um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	ID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if ID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

	if err = usuario.ValidarUsuario("edicao"); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)

	if err = repositorio.Atualizar(ID, usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, nil)
}

// Deleta um usuário do banco de dados
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
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

	if err = repositorio.Deletar(ID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
