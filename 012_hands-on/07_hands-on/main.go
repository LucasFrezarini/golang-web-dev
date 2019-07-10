package main

import (
	"log"
	"os"
	"text/template"
)

type food struct {
	Name        string
	Description string
	Price       float64
}

type meal struct {
	Name  string
	Foods []food
}

type restaurant struct {
	Restaurant string
	Menu       []meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menus := getMenus()

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}

func getMenus() []restaurant {
	return []restaurant{
		restaurant{
			Restaurant: "In the end of galaxy",
			Menu: []meal{
				meal{
					Name: "Lunch",
					Foods: []food{
						food{"Hamburger", "Very good", 34.40},
						food{"Parmeggiana", "The best food ever", 29.90},
					},
				},
				meal{
					Name: "Breakfast",
					Foods: []food{
						food{"Bread", "The good and old bread", 2.00},
					},
				},
				meal{
					Name: "Dinner",
					Foods: []food{
						food{"Some Fancy Pasta", "Expensive and fancy pasta", 42.00},
					},
				},
			},
		},
		restaurant{
			Restaurant: "Of hotel california",
			Menu: []meal{
				meal{
					Name: "Lunch",
					Foods: []food{
						food{"Hamburger", "Very good", 34.40},
						food{"Parmeggiana", "The best food ever", 29.90},
					},
				},
				meal{
					Name: "Breakfast",
					Foods: []food{
						food{"Bread", "The good and old bread", 2.00},
					},
				},
				meal{
					Name: "Dinner",
					Foods: []food{
						food{"Some Fancy Pasta", "Expensive and fancy pasta", 42.00},
					},
				},
			},
		},
	}
}
