package models

type CustomError struct {
	Mensagem string
	Erro     error
}

func (e CustomError) Error() string {
	return e.Mensagem + ": " + e.Erro.Error()
}
