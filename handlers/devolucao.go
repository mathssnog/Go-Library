/* Rota para devolução de livro:

   Será utilizado a mesma tabela de emprestimos para 'deletarmos' o livro
   que foi feito o empréstimo.

   A princípio vamos seguir dessa forma, e depois podemos fazer algumas
   alterações de refinamento

*/

package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Rota para deletar o livro que foi feito o empréstimo caso ele seja devolvido:

type RemoveEmprestimo struct {
	Membro string `json:"email_usuario"`
}

func Devolucao(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RemoveEmprestimo
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		email_usuario := req.Membro
		_, err := db.Exec("DELETE FROM emprestimo WHERE email_usuario=$1", email_usuario)
		if err != nil {
			fmt.Println("Error delete devolucao:", err) // Log para debbug
			return fmt.Errorf("could not delete emprestimo: %w", err)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User with email %s return successfully", email_usuario))
	}
}

//curl -X DELETE http://localhost:8080/devolucao/email_usuario -H "Content-Type: application/json" -d '{"email_usuario":"John dae"}'
