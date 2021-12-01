package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
)

type BookTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         string
}

func Book(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := BookTemplateVariables{
		Title:       "LiteratTeX",
		Pages:       root.NamePages,
		CurrentPage: "Livros",
		URL:         root.URL,
	}

	err := root.Templates.ExecuteTemplate(w, "book", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
