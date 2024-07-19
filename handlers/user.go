package handlers

import (
	"database/sql"
	"fmt"
	"go-library/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// São as funções que vão lidar com as rotas.

func AddUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(structs.User)
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
	}
}

// Função para listar usuários
// Lembrando que "db" é a conexão com o banco definido na função Connect()
// "rows" é o resultado da consulta 'Query'
// Gin REST api: https://deadsimplechat.com/blog/rest-api-with-golang-and-postgresql/

// Função para listar os usuários que estão no postgres
func ListUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT nome, email, id FROM usuario")
		if err != nil {
			return fmt.Errorf("could not get users: %w", err)
		}
		defer rows.Close() // Defer é utilizada para executar a função a qual foi chamada[rows.Close()], após a execução da função inteira em que ela se encontra. Essa função fecha a o resultado da consulta rows.

		var users []structs.User
		for rows.Next() { //loop nos dados retornados pela query
			var user structs.User
			if err := rows.Scan(&user.Name, &user.Email, &user.ID); err != nil {
				fmt.Println("Error:", err) // Log para debbug
				return fmt.Errorf("could not scan user: %w", err)
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("rows error: %w", err)
		}
		return c.JSON(http.StatusOK, users) // format type JSON
	}
}

type RemoveUser struct {
	ID int `json:"id"`
}

// Função para deletar usuários com base no seu ID:
func DeleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RemoveUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		id := req.ID
		_, err := db.Exec("DELETE FROM usuario WHERE id=$1", id)
		if err != nil {
			fmt.Println("Error inserting user:", err) // Log para debbug
			return fmt.Errorf("could not delete user: %w", err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User with ID %d deleted successfully", id))
	}
}
