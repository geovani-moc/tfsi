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
