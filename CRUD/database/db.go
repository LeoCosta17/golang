package database

import (
	"database/sql"
)

func ConnDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
}
