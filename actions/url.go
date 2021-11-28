package actions

import (
	"net/http"

	"github.com/geovani-moc/tfsi/controller"
)

func routes(app *App) {
	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.Home(w, r, &app.root)
	})

	app.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		cacheControlWrapper(http.FileServer(http.Dir("static")))))

	app.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.Code404(w, r, &app.root)
	})
}
