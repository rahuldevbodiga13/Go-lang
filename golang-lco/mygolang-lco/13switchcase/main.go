package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to dice game")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	var diceValue int = rand.Intn(6) + 1
	switch diceValue {
	case 1:
		fmt.Println("You can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll again")
	default:
		fmt.Println("What was that!")
	}

}
