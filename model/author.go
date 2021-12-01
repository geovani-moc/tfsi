package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
)

type AuthorTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         string
}

func Author(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := AuthorTemplateVariables{
		Title:       "Pesquisa sobre autores",
		Pages:       root.NamePages,
		CurrentPage: "Autores",
		URL:         root.URL,
	}

	err := root.Templates.ExecuteTemplate(w, "author", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
