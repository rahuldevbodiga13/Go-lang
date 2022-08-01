package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var dreamList [4]rune
	fmt.Println(dreamList)
	fmt.Println(len(dreamList))

	var someList = [4]string{"dfs"}
	fmt.Println(someList)

	fmt.Printf("Type of dreamlist is %T", someList)

}
