package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
}
