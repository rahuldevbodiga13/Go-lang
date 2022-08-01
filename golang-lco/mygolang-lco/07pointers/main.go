package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to pointers!!")
	var ptr *int
	fmt.Println("value of ptr is ", ptr)

	num := 27
	var refPtr = &num
	fmt.Println("Value of ref ppinter ", refPtr)
	fmt.Println("Value of *ref ppinter ", *refPtr)

}
