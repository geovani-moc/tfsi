package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
		routes(_app)
	}

	return _app
}

// NewApp Criação de intancia de um novo app
func NewApp() *App {

	app := &App{
		router:  mux.NewRouter(),
		version: "1.0.0",
		root: util.Root{
			Port:      ":3000",
			SiteName:  "LiteraTeX",
			URL:       "http://0.0.0.0:3000",
			Templates: parseTemplates(),
			NamePages: []string{
				"home",
				"Livros",
				"Autores",
			},
		},
	}
	return app
}

//Run inicia o servidor do app
func (app *App) Run() error {
	fmt.Println("LiteraTeX", app.version)
	fmt.Println("Link:" + app.root.URL)

	return http.ListenAndServe(app.root.Port, app.router)
}

func cacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=60") // em segundos
		h.ServeHTTP(w, r)
	})
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./view", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Print("erro ao realizar o parser do template, ", err)
			}
		}
		return err
	})

	if err != nil {
		log.Print("erro na verificação de diretorios do template, ", err)
	}

	return templ
}
