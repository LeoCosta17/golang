package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	resultados, err := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc
	`,
		usuarioID, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer resultados.Close()

	var publicacoes []models.Publicacao

	for resultados.Next() {
		var publicacao models.Publicacao
		if err = resultados.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	resultado, err := repositorio.db.Query(
		`SELECT p.*, u.nick FROM
		publicacoes p INNER JOIN usuarios u
		on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer resultado.Close()

	var publicacao models.Publicacao

	if resultado.Next() {
		if err = resultado.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return models.Publicacao{}, nil
		}
	}

	return publicacao, nil

}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, err := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if err != nil {
		return nil
	}
	defer statement.Close()

	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
