package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {

	statement, err := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha)VALUES (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repositorio Usuarios) Buscar(identificador string) ([]models.Usuario, error) {

	identificador = fmt.Sprintf("%%%s%%", identificador)

	results, err := repositorio.db.Query("SELECT id, nome, nick, email, dt_criado FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		identificador,
		identificador,
	)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var usuarios []models.Usuario

	for results.Next() {
		var usuario models.Usuario

		if err = results.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.DataCriado); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {

	result, err := repositorio.db.Query(
		"SELECT id, nome, nick, email, dt_criado FROM usuarios WHERE id = ?",
		ID,
	)

	if err != nil {
		return models.Usuario{}, err
	}

	defer result.Close()

	var Usuario models.Usuario

	if result.Next() {
		if err = result.Scan(
			&Usuario.ID,
			&Usuario.Nome,
			&Usuario.Nick,
			&Usuario.Email,
			&Usuario.DataCriado,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return Usuario, nil
}

func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return nil
	}
	defer statement.Close()

	fmt.Println("Dados recebidos por repositorio Atualizar: ", usuario)

	_, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil

}
