package routes

import "net/http"

// Representa a estrutura básica de uma rota no sistema
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}
