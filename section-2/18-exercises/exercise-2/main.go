package main

import (
	"log"
	"os"
	"text/template"
)

type region int

func (r region) String() string {
	switch r {
	case Southern:
		return "Southern"
	case Central:
		return "Central"
	case Northern:
		return "Northern"
	}
	return "Unknown"
}

const (
	Southern region = iota
	Central
	Northern
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	hotels := []hotel{
		hotel{"Hotel California", "Lovely Pl", "Los Angeles", "12345", Southern},
		hotel{"Hotel Palifornia", "Lovely Pl", "San Francisco", "12345", Central},
		hotel{"Hotel Dalifornia", "Lovely Pl", "San Diego", "12345", Northern},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
