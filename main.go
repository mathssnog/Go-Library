package main

// Precisamos instalar o echo: go get -u github.com/labstack/echo/v4
// Após o go get, inicializei go mod:  go mod ini go-library
// Depois disso eu realizei o comando: go mod tidy
// garantindo assim que todas as dependências do módulo estejam atualizadas e corretamente refletidas no arquivo go.mod

// Reference links:
/*
	- https://dev.to/janirefdez/connect-rest-api-to-database-with-go-d8m - 07/06/2024
	- https://hevodata.com/learn/golang-postgres/
*/

import (
	"go-library/database"
	"go-library/routes"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

// Função que contém as informações para a conexão com o database 'livros'
// 'sql.Open' conecta ao database, e posteriormente vamos checar se esta conexão foi bem sucedida.

func main() {
	e := echo.New()
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer database.CloseConnection(db)

	routes.Routes(e, db)

	// Incia o servidor em todas as interfaces (0.0.0.0)
	go func() {
		if err := e.Start("0.0.0.0:8080"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Block the main goroutine to keep the server running for dockerfile/docker-compose
	select {}
}

//Após a criação da rota para usuário/livro novo, precisamos armazenar essa requisição no banco de dados Postgres.
//Quando conseguirmos realizar o armazenamento, podemos criar as 2 rotas subsequentes que são para mostrar os usuários e livros cadastrados.
