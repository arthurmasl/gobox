package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Pizza struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	c := colly.NewCollector()

	var list []Pizza

	c.OnHTML(".prodcut-title", func(e *colly.HTMLElement) {
		pizza := Pizza{Name: e.ChildText("h3"), Description: strings.TrimSpace(e.DOM.Next().Text())}

		list = append(list, pizza)
	})

	c.OnScraped(func(r *colly.Response) {
		pizzaData, mErr := json.MarshalIndent(list, "", "  ")
		if mErr != nil {
			fmt.Println(mErr)
		}

		writeErr := os.WriteFile("pizza.json", pizzaData, 0644)
		if writeErr != nil {
			fmt.Println(writeErr)
		}
	})

	c.Visit("https://doublepizza.lt/produktai/picos")
}
