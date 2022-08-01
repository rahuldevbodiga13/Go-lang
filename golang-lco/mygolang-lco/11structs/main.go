package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang !!")

	rahana := User{"rahana", "rahana@go.dev", true, 22}
	fmt.Printf("User rahana info : %+v\n", rahana)
	fmt.Println("Name is ", rahana.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
