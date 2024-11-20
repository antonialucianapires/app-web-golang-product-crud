package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaDB() *sql.DB {
	conexao := "user=user dbname=gestao_produto password=user host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}