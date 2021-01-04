package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	users := []user{
		user{
			Name:  "Buddha",
			Motto: "The belief of no beliefs",
			Admin: false,
		},
		user{
			Name:  "Gandhi",
			Motto: "Be the change",
			Admin: true,
		},
		user{
			Name:  "",
			Motto: "Nobody",
			Admin: true,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
