package main

import "fmt"

func main() {
	defer fmt.Println("Welcome to ")
	fmt.Println("defer in golang")
	defer fmt.Println("this is defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
