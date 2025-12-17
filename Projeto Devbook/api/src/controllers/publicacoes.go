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

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	var publicacao models.Publicacao

	if err := json.NewDecoder(r.Body).Decode(&publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID

	if err = publicacao.ValidarPublicacao(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioPublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacoes, err := repositorio.Buscar(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	publicacaoID, err := strconv.ParseUint(r.PathValue("publicacao_id"), 10, 64)
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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)
	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusFound, publicacao)

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	publicacaoID, err := strconv.ParseUint(r.PathValue("publicacao_id"), 10, 64)
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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacaoSalvaBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	var publicacao models.Publicacao

	if err = json.NewDecoder(r.Body).Decode(&publicacao); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = publicacao.ValidarPublicacao(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorio.Atualizar(publicacaoID, publicacao); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	publicacaoID, err := strconv.ParseUint(r.PathValue("publicacao_id"), 10, 64)
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

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacaoSalvaBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível excluir uma publicação que não seja sua"))
		return
	}

	if err = repositorio.Deletar(publicacaoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := strconv.ParseUint(r.PathValue("usuario_id"), 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.DBConn()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioPublicacoes(db)

	publicacoes, err := repositorio.BuscarPorUsuario(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}
