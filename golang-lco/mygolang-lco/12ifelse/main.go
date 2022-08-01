package main

import "fmt"

func main() {
	fmt.Println("If-Else in Golang!!")
	loginCount := 27
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount >= 10 && loginCount < 27 {
		result = "paid user"
	} else {
		result = "our user"
	}

	fmt.Println(result)

	if num := 3; num < 3 {
		fmt.Println("num is less than 3")
	} else {
		fmt.Println("num is greater than or equals 3")
	}

}
