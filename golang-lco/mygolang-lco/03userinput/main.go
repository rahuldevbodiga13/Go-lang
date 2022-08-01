package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, size, err := reader.ReadRune()
	fmt.Println("thanks for giving error input", input, size, err)
}
