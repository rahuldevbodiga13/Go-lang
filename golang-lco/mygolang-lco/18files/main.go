package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This should go inside another file - Rahana"
	var file, err = os.Create("./log.txt")

	CheckNilError(err)

	length, err := io.WriteString(file, content)
	CheckNilError(err)
	fmt.Println("Length is ", length)
	defer file.Close()
	ReadFile("./log.txt")

}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	CheckNilError(err)
	fmt.Println("Data inside "+filename+" is ", string(databyte))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
