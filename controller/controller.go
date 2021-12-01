package controller

import (
	"net/http"

	"github.com/geovani-moc/tfsi/model"
	"github.com/geovani-moc/tfsi/util"
)

//Home initial page controller
func Home(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Home(w, r, root)
}

func Book(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Book(w, r, root)
}

func Author(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Author(w, r, root)
}

//Code404 redirect to page not foud
func Code404(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Code404(w, r, root)
}
