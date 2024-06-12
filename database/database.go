package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

// Função que contém as informações para a conexão com o database 'livros'
// 'sql.Open' conecta ao database, e posteriormente vamos checar se esta conexão foi bem sucedida.

func Connect() (*sql.DB, error) {
	connectInfo := "host=localhost user=livros_admin password=123 dbname=livros port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database conectado!!")
	return db, nil
}

// Função que termina a conexão com o database 'livros'
func CloseConnection(db *sql.DB) {
	defer db.Close()
}
