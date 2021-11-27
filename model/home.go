package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
)

//HomeTemplateVariables variables
type HomeTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         string
}

//Home página inicial do sistema
func Home(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variables := HomeTemplateVariables{
		Title:       "Funções disponiveis",
		Pages:       []string{"Buscar aurtor", "obras de um autor"},
		CurrentPage: "home",
		URL:         root.URL,
	}

	err := root.Templates.ExecuteTemplate(w, "home", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
