package database

import (
	"database/sql"
	"fmt"
	"log"
	"slices"

	_ "github.com/mattn/go-sqlite3"
)

func ConnDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func init() {
	fmt.Println("Verificando existências tabela do banco de dados...")

	db, err := ConnDatabase()
	if err != nil {
		return
	}

	resultados, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return
	}
	defer resultados.Close()

	var tabelas []string

	for resultados.Next() {
		var tabela string
		if err := resultados.Scan(&tabela); err != nil {
			return
		}

		tabelas = append(tabelas, tabela)
	}

	tabelaContato := slices.Contains(tabelas, "contatos")

	if !tabelaContato {
		statement := `CREATE TABLE contatos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT NOT NULL UNIQUE,
			email TEXT,
			telefone TEXT
		);`

		_, err = db.Exec(statement)
		if err != nil {
			log.Println("Erro ao criar tabelas do banco de dados!", err)
			return
		}

		fmt.Println("Tabela de contatos criada com sucesso!")
	}

	fmt.Println("Tabelas verificadas e criadas com sucesso!")
}
