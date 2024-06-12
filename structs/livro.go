package structs

type Livro struct {
	Titulo string `json:"titulo" xml:"titulo" form:"titulo" query:"titulo"`
	Autor  string `json:"autor" xml:"autor" form:"autor" query:"autor"`
	Year   int    `json:"year" xml:"year" form:"year" query:"year"`
	Editor string `json:"editor" xml:"editor" form:"editor" query:"editor"`
}
