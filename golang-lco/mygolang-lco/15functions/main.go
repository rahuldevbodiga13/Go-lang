package main

import "fmt"

func greet() {
	fmt.Println("Hello!!")
}

func main() {
	fmt.Println("Welcome to functions")
	greet()
	result := adder(2, 5)
	fmt.Println("The result is ", result)

	proRes, _, _, _ := proAdder(1, 2, 4, 65, 5, 5)
	fmt.Println("Pro result is ", proRes)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(varargs ...int) (int, string, string, int) {
	sum := 0
	for i := range varargs {
		sum += varargs[i]
	}
	return sum, "hols", "df", 64
}
