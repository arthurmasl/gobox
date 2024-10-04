package main

import (
	"fmt"
	"log"

	"gobox/pkg/utils"
)

func main() {
	log.SetFlags(log.Lshortfile)

	sum, err := utils.Summy(1, 1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sum)

	sum2, err := utils.Summy(555, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(sum2)

	sum3, err := utils.Summy(2, 2)
	if err != nil {
		return
	}

	fmt.Println(sum3)
}
