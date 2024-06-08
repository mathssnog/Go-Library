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
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"` // Podemos receber dados como json, xml, form, query
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
}

type Livro struct {
	Titulo string `json:"titulo" xml:"titulo" form:"titulo" query:"titulo"`
	Autor  string `json:"autor" xml:"autor" form:"autor" query:"autor"`
	Year   int    `json:"year" xml:"year" form:"year" query:"year"`
	Editor string `json:"editor" xml:"editor" form:"editor" query:"editor"`
}

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

func main() {
	e := echo.New()
	db, err := Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao database!!: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Worldaaaaaa!")
	})

	e.POST("add/user", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		//Dados do header
		fmt.Println("Headers da requisição:")
		for name, values := range c.Request().Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Registrar usuário no banco de dados 'livros' na tabela 'usuario'
		_, err := db.Exec("INSERT INTO usuario (id, nome, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
		if err != nil {
			fmt.Println("Error inserting user:", err) // Log para debbug
			return fmt.Errorf("could not insert user: %w", err)
		}

		// Impressão dos dados do User
		fmt.Printf("Received User: Name=%s, Email=%s, Id=%s\n", user.Name, user.Email, user.ID)
		return c.JSON(http.StatusOK, user)
	})

	e.POST("add/livro", func(c echo.Context) error {
		livro := new(Livro)
		if err := c.Bind(livro); err != nil {
			return err
		}

		//Dados do header
		fmt.Println("Headers da requisição:")
		for name, values := range c.Request().Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Registrando o livro na table 'livro'
		_, err := db.Exec("INSERT INTO livro (titulo, autor, year, editor) VALUES ($1, $2, $3, $4)", livro.Titulo, livro.Autor, livro.Year, livro.Editor)
		if err != nil {
			fmt.Println("Error inserting user:", err) // Log para debbug
			return fmt.Errorf("could not insert livro: %w", err)
		}

		// Impressão dos dados do User
		fmt.Printf("Received livro: Titulo=%s, Autor=%s, Editor=%s, Year=%d\n", livro.Titulo, livro.Autor, livro.Editor, livro.Year)
		return c.JSON(http.StatusOK, livro)
	})

	// Rota para deletar usuários com base no seu ID:
	e.DELETE("delete/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM usuario WHERE id=$1", id)
		if err != nil {
			fmt.Println("Error inserting user:", err) // Log para debbug
			return fmt.Errorf("could not delete user: %w", err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User with ID %s deleted successfully", id))
	})

	// Rota para deletar livros com base no Título:
	e.DELETE("delete/livro/titulo", func(c echo.Context) error {
		titulo := c.Param("titulo")
		_, err := db.Exec("DELETE FROM livro WHERE titulo=$1", titulo)
		if err != nil {
			fmt.Println("Error inserting user:", err) // Log para debbug
			return fmt.Errorf("could not delete user: %w", err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User with Título %s deleted successfully", titulo))
	})

	// Rota para listar usuários
	// Lembrando que "db" é a conexão com o banco definido na função Connect()
	// "rows" é o resultado da consulta 'Query'
	// Gin REST api: https://deadsimplechat.com/blog/rest-api-with-golang-and-postgresql/
	e.GET("list/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT nome, email, id FROM usuario")
		if err != nil {
			return fmt.Errorf("could not get users: %w", err)
		}
		defer rows.Close() // Defer é utilizada para executar a função a qual foi chamada, nesse caso foi o rows.Close(). Essa função fecha a o resultado da consulta rows.

		var users []User
		for rows.Next() { //loop nos dados retornados pela query
			var user User
			if err := rows.Scan(&user.Name, &user.Email, &user.ID); err != nil {
				fmt.Println("Error inserting user:", err) // Log para debbug
				return fmt.Errorf("could not scan user: %w", err)
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("rows error: %w", err)
		}
		return c.JSON(http.StatusOK, users) // format type JSON
	})

	// Rota para listar livros
	e.GET("list/livro", func(c echo.Context) error {
		rows, err := db.Query("SELECT titulo, autor, year, editor FROM livro")
		if err != nil {
			return fmt.Errorf("could not get livro: %w", err)
		}
		defer rows.Close()

		var livros []Livro
		for rows.Next() {
			var livro Livro
			if err := rows.Scan(&livro.Titulo, &livro.Autor, &livro.Year, &livro.Editor); err != nil {
				return fmt.Errorf("could not scan livro: %w", err)
			}
			livros = append(livros, livro)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("rows error: %w", err)
		}
		return c.JSON(http.StatusOK, livros)
	})

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
