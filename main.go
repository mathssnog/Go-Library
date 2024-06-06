package main

// Precisamos instalar o echo: go get -u github.com/labstack/echo/v4
// Após o go get, inicializei go mod:  go mod ini go-library
// Depois disso eu realizei o comando: go mod tidy
// garantindo assim que todas as dependências do módulo estejam atualizadas e corretamente refletidas no arquivo go.mod

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
}

type Book struct {
	Titulo string `json:"titulo" xml:"titulo" form:"titulo" query:"titulo"`
	Autor  string `json:"autor" xml:"autor" form:"autor" query:"autor"`
	Year   int    `json:"year" xml:"year" form:"year" query:"year"`
	Editor string `json:"editor" xml:"editor" form:"editor" query:"editor"`
}

func main() {
	e := echo.New()

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

		// Impressão dos dados do User
		fmt.Printf("Received User: Name=%s, Email=%s, Id=%s\n", user.Name, user.Email, user.ID)
		return c.JSON(http.StatusOK, user)
	})

	e.POST("add/book", func(c echo.Context) error {
		book := new(Book)
		if err := c.Bind(book); err != nil {
			return err
		}

		//Dados do header
		fmt.Println("Headers da requisição:")
		for name, values := range c.Request().Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Impressão dos dados do User
		fmt.Printf("Received Book: Titulo=%s, Autor=%s, Editor=%s, Year=%d\n", book.Titulo, book.Autor, book.Editor, book.Year)
		return c.JSON(http.StatusOK, book)
	})

	// Incia o servidor em todas as interfaces (0.0.0.0)
	go func() {
		if err := e.Start("0.0.0.0:8080"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Block the main goroutine to keep the server running
	select {}
}

//Após a criação da rota para usuário/livro novo, precisamos armazenar essa requisição no banco de dados Postgres.
//Quando conseguirmos realizar o armazenamento, podemos criar as 2 rotas subsequentes que são para mostrar os usuários e livros cadastrados.
