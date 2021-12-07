package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
)

//Code404TemplateVariables variaveis da pagina
type Code404TemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         string
}

//Code404 generate page
func Code404(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := Code404TemplateVariables{
		Title:       "Página não encontrada",
		Pages:       root.NamePages,
		CurrentPage: "httpStatusCode",
		URL:         r.Host,
	}

	err := root.Templates.ExecuteTemplate(w, "not-found", variables)
	if nil != err {
		log.Println("Error ao gerar pagina nao encontrada, ", err)
	}
}
