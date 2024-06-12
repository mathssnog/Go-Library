package handlers

import (
	"database/sql"
	"fmt"
	"go-library/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// São as funções que vão lidar com as rotas.

func Addlivro(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		livro := new(structs.Livro)
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
			fmt.Println("Error inserting livro:", err) // Log para debbug
			return fmt.Errorf("could not insert livro: %w", err)
		}

		// Impressão dos dados do User
		fmt.Printf("Received livro: Titulo=%s, Autor=%s, Editor=%s, Year=%d\n", livro.Titulo, livro.Autor, livro.Editor, livro.Year)
		return c.JSON(http.StatusOK, livro)
	}
}

func ListLivro(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT titulo, autor, year, editor FROM livro")
		if err != nil {
			return fmt.Errorf("could not get livro: %w", err)
		}
		defer rows.Close()

		var livros []structs.Livro
		for rows.Next() {
			var livro structs.Livro
			if err := rows.Scan(&livro.Titulo, &livro.Autor, &livro.Year, &livro.Editor); err != nil {
				return fmt.Errorf("could not scan livro: %w", err)
			}
			livros = append(livros, livro)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("rows error: %w", err)
		}
		return c.JSON(http.StatusOK, livros)
	}
}

func Deletelivro(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		titulo := c.Param("titulo")
		_, err := db.Exec("DELETE FROM livro WHERE titulo=$1", titulo)
		if err != nil {
			fmt.Println("Error inserting livro:", err) // Log para debbug
			return fmt.Errorf("could not delete livro: %w", err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User with Título %s deleted successfully", titulo))
	}
}
