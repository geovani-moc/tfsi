package model

import (
	"html/template"
	"log"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
)

//HomeTemplateVariables variables
type HomeTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         template.URL
}

//Home página inicial do sistema
func Home(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variables := HomeTemplateVariables{
		Title:       "LiteraTeX",
		Pages:       root.NamePages,
		CurrentPage: "home",
		URL:         template.URL(r.Host),
	}

	err := root.Templates.ExecuteTemplate(w, "home", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
