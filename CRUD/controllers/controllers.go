package controllers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}
