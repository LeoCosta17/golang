package controllers

import "net/http"

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela login!"))
}
