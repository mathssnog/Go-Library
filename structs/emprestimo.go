package structs

import "time"

type Emprestimo struct {
	Membro     string    `json:"membro" xml:"membro" form:"membro" query:"membro"`
	Nome_Livro string    `json:"nome_livro" xml:"nome_livro" form:"nome_livro" query:"nome_livro"`
	Data       time.Time `json:"data" xml:"data" form:"data" query:"data"`
}

/* Exemplo de requisição POST:

curl -X POST http://localhost:8080/emprestimo \
-H "Content-Type: application/json" \
-d '{"membro": "John Doe", "nome_livro": "1989", "data": "1999-01-01T00:00:00Z"}'

*/
