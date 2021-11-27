package util

import (
	"html/template"
)

type Root struct {
	Templates *template.Template
	Port      string
	SiteName  string
	URL       string
}
