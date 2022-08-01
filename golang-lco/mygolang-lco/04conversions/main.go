package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Rahana Store")
	fmt.Println("Please provide us with a rating b/w 1 - 5 :)")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving us a rating of " + rating)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating : ", newRating+1)
	}

}
