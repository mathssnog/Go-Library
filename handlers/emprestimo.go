/* Rota para empréstimo de livro:

- Verifica se temos o livro cadastro na tabela de livros;
- Verifica o número de cópias cadastradas desse livro;

[Ajustar: adicionar coluna de "copias" na tabela de livros
		  Caso for cadastrar livros, verificar se existe,
		  se positivo, incrementar a coluna de "copias"]
*/

/* Criar tabelas de empréstimo:

- Caso as condições de empréstimo forem satisfeitas,
será incluído na tabela "emprestimo":
	- email_usuario;
	- nome_livro;
	- data_emprestimo;
CREATE TABLE emprestimo (
    email_usuario VARCHAR(255) NOT NULL,
    nome_livro VARCHAR(255) NOT NULL,
    data_emprestimo DATE NOT NULL,
    PRIMARY KEY (email_usuario, nome_livro, data_emprestimo)
);
##Lembrando que o comando PRIMARY KEY garante que a combinação
(email_usuario, nome_livro, data_emprestimo) --  seja única para
cada registro!!
*/

package handlers

import (
	"database/sql"
	"fmt"
	"go-library/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// São as funções que vão lidar com as rotas.

func AddEmprestimo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		emprestimo := new(structs.Emprestimo)
		if err := c.Bind(emprestimo); err != nil {
			return err
		}
		//Dados do header
		fmt.Println("Headers da requisição:")
		for name, values := range c.Request().Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Registrando o emprestimo na table 'emprestimo'
		_, err := db.Exec("INSERT INTO emprestimo (email_usuario, nome_livro, data_emprestimo) VALUES ($1, $2, $3)", emprestimo.Membro, emprestimo.Nome_Livro, emprestimo.Data)
		if err != nil {
			fmt.Println("Error inserting emprestimo:", err) // Log para debbug
			return fmt.Errorf("could not insert emprestimo: %w", err)
		}

		// Impressão dos dados do User
		//		fmt.Printf("Received emprestimo: Membro=%s, Nome_Livro=%s, data=%s, Year=%d\n", livro.Titulo, livro.Autor, livro.Editor, livro.Year)
		return c.JSON(http.StatusOK, emprestimo)
	}
}

func ListEmprestimo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT email_usuario, nome_livro, data_emprestimo FROM emprestimo")
		if err != nil {
			return fmt.Errorf("could not get emprestimo: %w", err)
		}
		defer rows.Close()

		var emprestimos []structs.Emprestimo
		for rows.Next() {
			var emprestimo structs.Emprestimo
			if err := rows.Scan(&emprestimo.Membro, &emprestimo.Nome_Livro, &emprestimo.Data); err != nil {
				return fmt.Errorf("could not scan emprestimo: %w", err)
			}
			emprestimos = append(emprestimos, emprestimo)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("rows error: %w", err)
		}
		return c.JSON(http.StatusOK, emprestimos)
	}
}
