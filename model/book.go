package model

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/geovani-moc/tfsi/entity"
	"github.com/geovani-moc/tfsi/util"
)

type BookTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
	URL         string
	Op          string
	Search      string
	InfoBook    entity.InfoBook
	EditionBook entity.EditionBook
}

func Book(w http.ResponseWriter, r *http.Request, root *util.Root) {
	errInternal := ""
	variables := BookTemplateVariables{
		Title:       "Pesquisa sobre livros",
		Pages:       root.NamePages,
		CurrentPage: "Livros",
		URL:         root.URL,
		Search:      "",
	}

	if r.Method != http.MethodPost {
		err := root.Templates.ExecuteTemplate(w, "book", variables)
		if err != nil {
			log.Print("Template executing error: ", err)
			internalError(w, root)
		}
		return
	}

	variables.Search = r.FormValue("buscar")
	variables.Op = r.FormValue("op")

	if variables.Search == "" {

		err := root.Templates.ExecuteTemplate(w, "internalError", variables)
		if err != nil {
			log.Println("Template executing error: ", err)
			internalError(w, root)
			return
		}
		return
	}
	if variables.Op == "1" {
		variables.InfoBook, errInternal = searchInfoBook(variables.Search)
	} else if variables.Op == "2" {
		variables.EditionBook, errInternal = searchEditionBookByCode(variables.Search)
	} else if variables.Op == "3" {
		variables.EditionBook, errInternal = searchEditionBookByISBN(variables.Search)
	}

	if errInternal != "" {
		internalError(w, root)
		log.Println(errInternal)
		return
	}

	err := root.Templates.ExecuteTemplate(w, "book", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}

func searchInfoBook(code string) (entity.InfoBook, string) {
	var result entity.InfoBook
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://openlibrary.org/works/" + code + ".json"

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

func searchEditionBookByCode(code string) (entity.EditionBook, string) {
	var result entity.EditionBook
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://openlibrary.org/books/" + code + ".json"

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

func searchEditionBookByISBN(code string) (entity.EditionBook, string) {
	var result entity.EditionBook
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://openlibrary.org/isbn/" + code + ".json"

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
