package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"})
	if err != nil {
		log.Fatalln(err)
	}
}
