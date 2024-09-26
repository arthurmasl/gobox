package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type PizzaJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Pizza struct {
	name        string
	ingridients []string
}

func main() {
	b, rErr := os.ReadFile("pizza.json")
	if rErr != nil {
		fmt.Println(rErr)
	}

	list := []PizzaJson{}
	uErr := json.Unmarshal(b, &list)
	if uErr != nil {
		fmt.Println(uErr)
	}

	pizzas := []Pizza{}
	ingridientNames := []string{}
	ingridients := make(map[string]int)

	for _, p := range list {
		ingridientsSlice := strings.Split(p.Description[:len(p.Description)-1], ", ")

		for _, ingridient := range ingridientsSlice {
			if _, exists := ingridients[ingridient]; exists {
				ingridients[ingridient]++
			} else {
				ingridients[ingridient] = 0
				ingridientNames = append(ingridientNames, ingridient)
			}
		}

		pizza := Pizza{
			name:        p.Name,
			ingridients: ingridientsSlice,
		}
		pizzas = append(pizzas, pizza)
	}

	// fmt.Printf("%#v\n", pizzas[0])
	// fmt.Printf("%#v\n", ingridients)
	// fmt.Println(ingridientNames)

	fmt.Printf("%#v\n", findPizzaNames(ingridientNames[6], &pizzas))
}

func findPizzaNames(ingridient string, pizzas *[]Pizza) []string {
	fmt.Println(ingridient)
	list := []string{}

	for _, p := range *pizzas {
		for _, i := range p.ingridients {
			if ingridient == i {
				list = append(list, p.name)
			}
		}
	}

	return list
}

func findPizzas(ingridient string, pizzas *[]Pizza) []*Pizza {
	list := []*Pizza{}

	for _, p := range *pizzas {
		for _, i := range p.ingridients {
			if ingridient == i {
				list = append(list, &p)
			}
		}
	}

	return list
}
