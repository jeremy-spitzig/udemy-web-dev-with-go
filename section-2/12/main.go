package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	sages := []sage{
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
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
