package structs

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"` // Podemos receber dados como json, xml, form, query
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
}
