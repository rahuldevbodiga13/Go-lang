package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(-1) + 1)
	// }

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(6) + 1)

	randId, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randId)

}
