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

	// authorName := r.URL.Query().Get("author_name")
	// if authorName != "" {
	// 	//fazer a busca automaticamente
	// }

	//pegar nome no formulario
	result := searchAuthorByName("j k rowling")
	print(result.NumFound)

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
