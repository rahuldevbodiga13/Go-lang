package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("Web request handling in GOlang")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response is ", string(databytes))

}
