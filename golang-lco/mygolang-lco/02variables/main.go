package main

import (
	"fmt"
)

const JwtToken string = "sdgkjnfjnfjdnkfjs" //public

func main() {
	var username string = "rahana"
	fmt.Println(username)
	fmt.Printf("variable %v is of type : %T \n", username, username)

	var isMarried bool = true
	fmt.Println(isMarried)
	fmt.Printf("variable is of type : %T \n", isMarried)

	var smallVal int64 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.456789465768767877
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	var anotherVar string
	fmt.Println(anotherVar)

	anotherType := 10
	fmt.Println(anotherType)

	fmt.Println(JwtToken)

}
