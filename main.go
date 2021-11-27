package main

import (
	"log"

	"github.com/geovani-moc/tfsi/actions"
)

func main() {
	app := actions.BuildApp()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
