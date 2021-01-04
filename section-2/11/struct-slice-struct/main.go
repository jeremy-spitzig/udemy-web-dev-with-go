package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	data := items{
		Wisdom: []sage{
			sage{
				Name:  "Buddha",
				Motto: "The belief of no beliefs",
			},
			sage{
				Name:  "Gandhi",
				Motto: "Be the change",
			},
			sage{
				Name:  "MLK",
				Motto: "Hatred never ceases with hatred but with love alone is healed",
			},
			sage{
				Name:  "Jesus",
				Motto: "Love all",
			},
			sage{
				Name:  "Muhammad",
				Motto: "To overcome evil with good is good, to resist evil by evil is evil",
			},
		},
		Transport: []car{
			car{
				Manufacturer: "Ford",
				Model:        "F150",
				Doors:        2,
			},
			car{
				Manufacturer: "Toyota",
				Model:        "Corolla",
				Doors:        4,
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
