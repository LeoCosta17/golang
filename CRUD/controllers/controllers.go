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
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
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
		http.Error(w, "Erro ao criar contato", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
