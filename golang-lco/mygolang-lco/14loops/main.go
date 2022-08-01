package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"mon", "teus", "wed", "thurs", "sat", "sun"}
	fmt.Println("Days are ", days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("index %d value is %v\n", index, value)
	}

	arbitraryValue := 1

	for arbitraryValue < 27 {
		if arbitraryValue == 2 {
			goto rahana
		}
		if arbitraryValue == 5 {
			break
		}
		fmt.Println("The value is ", arbitraryValue)
		arbitraryValue++
	}

rahana:
	fmt.Println("We're hereee")
}
