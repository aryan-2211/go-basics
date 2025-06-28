package main

import "fmt"

func main() {
	greeter()

	result := adder(4, 5)
	fmt.Println(result)

	proRes, message := proAdder(2, 3, 4, 5, 6, 7)
	fmt.Println(proRes)
	fmt.Println(message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func greeter() {
	fmt.Println("Namastey from golang")
}

//proFunc
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "default message"
}
