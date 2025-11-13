package controllers

import (
	"crud/models"
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func CriarContatoHandler(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusCreated)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
