package main

import (
	"log"
	"os"
	"text/template"
)

type menuItemType int

const (
	Food menuItemType = iota
	Beverage
)

func (mit menuItemType) String() string {
	switch mit {
	case Food:
		return "Food"
	case Beverage:
		return "Beverage"
	}
	return "Unknown"
}

// 1. Create a data structure to pass to a template which
// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
type menuItem struct {
	Name        string
	Description string
	Price       float64
	Type        menuItemType
}

type menu struct {
	Breakfast []menuItem
	Lunch     []menuItem
	Dinner    []menuItem
}

type restaurant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	restaurants := []restaurant{
		restaurant{
			Name: "McBonalds",
			Menu: menu{
				Breakfast: []menuItem{
					menuItem{"Eggs", "Eggy", 5.00, Food},
					menuItem{"Pancakes", "Fluffy", 10.00, Food},
					menuItem{"Orange Juice", "Citrusy", 2.00, Beverage},
				},
				Lunch: []menuItem{
					menuItem{"Sandwich", "Layery", 5.00, Food},
					menuItem{"Salad", "Leafy", 10.00, Food},
					menuItem{"Milk", "Milky", 2.00, Beverage},
				},
				Dinner: []menuItem{
					menuItem{"Steak", "Steaky", 5.00, Food},
					menuItem{"Fries", "Frenchy", 10.00, Food},
					menuItem{"Beer", "Hoppy", 2.00, Beverage},
				},
			},
		},
		restaurant{
			Name: "Vendy's",
			Menu: menu{
				Breakfast: []menuItem{
					menuItem{"Eggs", "Eggy", 5.00, Food},
					menuItem{"Pancakes", "Fluffy", 10.00, Food},
					menuItem{"Orange Juice", "Citrusy", 2.00, Beverage},
				},
				Lunch: []menuItem{
					menuItem{"Sandwich", "Layery", 5.00, Food},
					menuItem{"Salad", "Leafy", 10.00, Food},
					menuItem{"Milk", "Milky", 2.00, Beverage},
				},
				Dinner: []menuItem{
					menuItem{"Steak", "Steaky", 5.00, Food},
					menuItem{"Fries", "Frenchy", 10.00, Food},
					menuItem{"Beer", "Hoppy", 2.00, Beverage},
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
