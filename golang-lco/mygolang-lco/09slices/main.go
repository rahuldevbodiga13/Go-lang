package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	fruitList := []string{}
	fmt.Println(fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana", "Apple")
	fmt.Println("New Elements are ", fruitList)
	fruitList = append(fruitList, fruitList...)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 10
	highScores[1] = 654
	highScores[2] = 32
	highScores[3] = 27
	highScores = append(highScores, 3458)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove an element from a slice

	courses := []string{"android", "kotlin", "java", "golang", "mongodb"}
	fmt.Println("The list of courses are ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Update list is ", courses)

}
