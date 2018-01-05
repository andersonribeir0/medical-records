package config

import (
	"html/template"
	"log"
)

var TPL *template.Template

func init() {
	log.Println("Configuring templates.")
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
	log.Println("Templates configured.")
}
