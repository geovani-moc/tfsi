package actions

import (
	"net/http"

	"github.com/geovani-moc/tfsi/controller"
)

func Routes(app *App) {
	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.Home(w, r, &app.root)
	})
}
