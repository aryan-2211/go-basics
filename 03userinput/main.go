package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter rating for the pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input)
}
