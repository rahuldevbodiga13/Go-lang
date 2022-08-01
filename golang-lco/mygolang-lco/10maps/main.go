package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	language := make(map[int]int, 4)
	language[1] = 27
	language[2] = 27
	language[3] = 1327
	language[2] = 543
	language[4] = 54
	language[5] = 76

	fmt.Println(language)
	fmt.Println(len(language))
	fmt.Println(language[2])

	delete(language, 2)
	fmt.Println(language)

	for key, _ := range language {
		fmt.Printf("The value for key %d is \n", key)
	}

}
