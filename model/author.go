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
	Title         string
	Pages         []string
	CurrentPage   string
	URL           string
	Op            string
	Search        string
	AuthorByName  entity.AuthorByName
	AuthorByCode  entity.AuthorByCode
	WorksByAuthor entity.WorksByAuthor
}

func Author(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := AuthorTemplateVariables{
		Title:       "Pesquisa sobre autores",
		Pages:       root.NamePages,
		CurrentPage: "Autores",
		URL:         r.Host,
		Search:      "",
	}

	errInternal := ""

	if r.Method != http.MethodPost {
		err := root.Templates.ExecuteTemplate(w, "author", variables)
		if err != nil {
			log.Print("Template executing error: ", err)
			internalError(w, r, root)
			return
		}
		return
	}

	variables.Search = r.FormValue("buscar")
	variables.Op = r.FormValue("op")

	if variables.Search == "" {
		err := root.Templates.ExecuteTemplate(w, "internalError", variables)
		if err != nil {
			log.Print("Template executing error: ", err)
			internalError(w, r, root)
			return
		}
		return
	}

	if variables.Op == "1" {
		variables.AuthorByName, errInternal = searchAuthorByName(variables.Search)
	} else if variables.Op == "2" {
		variables.AuthorByCode, errInternal = searchAuthorByCode(variables.Search)
	} else if variables.Op == "3" {
		variables.WorksByAuthor, errInternal = searchWorksByAuthor(variables.Search)
	}

	if errInternal != "" {
		internalError(w, r, root)
		return
	}

	err := root.Templates.ExecuteTemplate(w, "author", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}

}

func searchAuthorByName(name string) (entity.AuthorByName, string) {
	client := &http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(http.MethodGet, "https://openlibrary.org/search/authors.json", nil)
	var result entity.AuthorByName

	if nil != err {
		return result, "erro interno: requisisção não pode ser criada."
	}

	query := request.URL.Query()
	query.Add("q", name)

	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if nil != err {
		return result, "erro interno: requisição não pode ser realizada."
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)

	if nil != err {
		return result, "erro interno: não foi possivel decodificar a resposta."
	}

	return result, ""
}

func searchAuthorByCode(code string) (entity.AuthorByCode, string) {
	var result entity.AuthorByCode
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://openlibrary.org/authors/" + code + ".json"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if nil != err {
		return result, "erro interno: requisisção não pode ser criada."
	}
	response, err := client.Do(request)

	if nil != err {
		return result, "erro interno: requisição não pode ser realizada."
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)

	if nil != err {
		return result, "erro interno: não foi possivel decodificar a resposta."
	}

	return result, ""
}

func searchWorksByAuthor(code string) (entity.WorksByAuthor, string) {
	var result entity.WorksByAuthor
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://openlibrary.org/authors/" + code + "/works.json"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if nil != err {
		return result, "erro interno: requisisção não pode ser criada."
	}
	response, err := client.Do(request)

	if nil != err {
		return result, "erro interno: requisição não pode ser realizada."
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)

	if nil != err {
		log.Println(err)
		return result, "erro interno: não foi possivel decodificar a resposta."
	}

	return result, ""
}

func internalError(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := AuthorTemplateVariables{
		Title:       "Erro interno",
		Pages:       root.NamePages,
		CurrentPage: "Internal error",
		URL:         r.Host,
		Search:      "",
	}
	err := root.Templates.ExecuteTemplate(w, "internalError", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
