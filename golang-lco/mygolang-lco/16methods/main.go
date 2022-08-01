package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	rahana := User{"rahana", "rahana@go.dev", true, 22, 2}
	fmt.Printf("User rahana info : %+v\n", rahana)
	fmt.Println("Name is ", rahana.Name)
	rahana.GetName()

	fmt.Println(rahana.GetEmail())
	rahana.SetEmail("rd@dev.go")
	fmt.Println(rahana.GetEmail())

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (user User) GetName() {
	fmt.Println("User Name : ", user.Name)
}

func (user User) SetEmail(email string) {
	user.Email = email
}

func (user User) GetEmail() string {
	return user.Email
}
