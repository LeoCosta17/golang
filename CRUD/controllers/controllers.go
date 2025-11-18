package controllers

import (
	"crud/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) { // Carrega a página inicial
	contatos, err := models.BuscarContatos()

	if err != nil {
		erro := models.CustomError{
			Mensagem: "Erro ao buscar contatos cadastrados",
			Erro:     err,
		}
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	w.WriteHeader(http.StatusOK)

	templates.ExecuteTemplate(w, "index.html", contatos)
}

func CriarContatoHandler(w http.ResponseWriter, r *http.Request) { // Cria um novo contato
	if r.Method != http.MethodPost {
		erro := models.CustomError{
			Mensagem: "Método não permitido",
			Erro:     nil,
		}

		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	nome := r.FormValue("inputName")
	email := r.FormValue("inputEmail")
	telefone := r.FormValue("inputPhone")

	contato := &models.Contato{
		Nome:     nome,
		Email:    email,
		Telefone: telefone,
	}

	err := contato.CriarContato(contato)
	if err != nil {
		erro := models.CustomError{
			Mensagem: "Erro ao criar contato",
			Erro:     err,
		}

		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func BuscarContatosHandler(w http.ResponseWriter, r *http.Request) ([]models.Contato, error) {
	contatos, err := models.BuscarContatos()
	if err != nil {
		erro := models.CustomError{
			Mensagem: "Erro ao buscar contatos cadastrados",
			Erro:     err,
		}
		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return nil, err
	}

	w.WriteHeader(http.StatusOK)
	return contatos, nil
}

func BuscarContatoHandler(w http.ResponseWriter, r *http.Request) {

}

func EditarContatoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		erro := models.CustomError{
			Mensagem: "Método não permitido",
			Erro:     nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	var contato models.Contato

	contato.ID, _ = strconv.Atoi(r.FormValue("editId"))
	contato.Nome = r.FormValue("editName")
	contato.Email = r.FormValue("editEmail")
	contato.Telefone = r.FormValue("editPhone")

	err := contato.EditarContato(&contato)
	if err != nil {
		erro := models.CustomError{
			Mensagem: "Erro ao editar contato.",
			Erro:     err,
		}
		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ExcluirContatoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		erro := models.CustomError{
			Mensagem: "Método não permitido",
			Erro:     nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}

	ID, _ := strconv.Atoi(r.PathValue("id"))

	err := models.ExcluirContato(ID)

	if err != nil {
		erro := models.CustomError{
			Mensagem: "Erro ao excluir contato.",
			Erro:     err,
		}

		w.WriteHeader(http.StatusInternalServerError)
		templates.ExecuteTemplate(w, "erro.html", erro)
		return
	}
}
