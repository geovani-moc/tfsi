package actions

import (
	"fmt"
	"net/http"

	"github.com/geovani-moc/tfsi/util"
	"github.com/gorilla/mux"
)

type App struct {
	router  *mux.Router
	version string
	root    util.Root
}

var _app *App

func BuildApp() *App {
	if nil == _app {
		_app = NewApp()
		Routes(_app)
	}

	return _app
}

// NewApp create the app
func NewApp() *App {

	app := &App{
		router:  mux.NewRouter(),
		version: "0.3",
		root: util.Root{
			Port:      ":8080",
			SiteName:  "Livros ***",
			Templates: nil,
			URL:       "http://localhost:8080",
		},
	}
	return app
}

//Run starts the server
func (app *App) Run() error {
	fmt.Println("Gcommerce", app.version)
	fmt.Println("Link: http://localhost" + app.root.Port)

	return http.ListenAndServe(app.root.Port, app.router)
}
