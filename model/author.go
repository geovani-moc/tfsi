package model

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/geovani-moc/tfsi/entity"
	"github.com/geovani-moc/tfsi/util"
)

type AuthorTemplateVariables struct {
	Title        string
	Pages        []string
	CurrentPage  string
	URL          string
	Op           string
	Search       string
	authorByName entity.AuthorByName
}

func Author(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := AuthorTemplateVariables{
		Title:       "Pesquisa sobre autores",
		Pages:       root.NamePages,
		CurrentPage: "Autores",
		URL:         root.URL,
		Search:      "",
	}

	if r.Method != http.MethodPost {
		err := root.Templates.ExecuteTemplate(w, "author", variables)
		if err != nil {
			log.Print("Template executing error: ", err)
		}
		return
	}

	variables.Search = r.FormValue("buscar")
	variables.Op = r.FormValue("op")

	if variables.Search == "" {
		//mostrar erro interno op= "-1" include tela de erro
		err := root.Templates.ExecuteTemplate(w, "author", variables)
		if err != nil {
			log.Print("Template executing error: ", err)
		}
		return
	}

	if variables.Op == "1" {
		variables.authorByName = searchAuthorByName(variables.Search)
	}

	err := root.Templates.ExecuteTemplate(w, "author", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}

}

func searchAuthorByName(name string) entity.AuthorByName {
	client := &http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(http.MethodGet, "https://openlibrary.org/search/authors.json", nil)

	if nil != err {
		log.Fatal("erro interno: requisisção não pode ser criada.")
	}

	query := request.URL.Query()
	query.Add("q", name)

	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if nil != err {
		log.Fatal("erro interno: requisição não pode ser realizada.")
	}

	defer response.Body.Close()

	var result entity.AuthorByName

	err = json.NewDecoder(response.Body).Decode(&result)

	if nil != err {
		log.Fatal("erro interno: não foi possivel decodificar a resposta.")
	}

	return result
}
