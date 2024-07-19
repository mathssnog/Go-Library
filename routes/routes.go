package routes

import (
	"database/sql"
	"net/http"

	"go-library/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, db *sql.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Worldaaaaaa!")
	})

	// Rota para adicionar usuário
	e.POST("add/user", handlers.AddUser(db))
	// Rota para adicionar livro
	e.POST("add/livro", handlers.Addlivro(db))
	// Rota para adicionar emprestimo
	e.POST("add/emprestimo", handlers.AddEmprestimo(db))

	// Rota para deletar usuário com base no seu ID
	e.DELETE("delete/user/id", handlers.DeleteUser(db))
	// Rota para deletar livros com base no seu Título
	e.DELETE("delete/livro/titulo", handlers.Deletelivro(db))
	// Rota para devolução de livros com base no email do usuário
	e.DELETE("devolucao/email_usuario", handlers.Devolucao(db))

	// Rota para listar os usuários no postgres
	e.GET("list/user", handlers.ListUser(db))
	// Rota para listar os livros no postgres
	e.GET("list/livro", handlers.ListLivro(db))
	// Rota para listar os livros no postgres
	e.GET("list/emprestimo", handlers.ListEmprestimo(db))
}
