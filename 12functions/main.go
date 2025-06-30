package main

import "fmt"

func main() {
	greeter()

	result := adder(4, 5)
	fmt.Println(result)

	proRes, message := proAdder(2, 3, 4, 5, 6, 7)
	fmt.Println(proRes)
	fmt.Println(message)

	result = func_name(5, 6)
	fmt.Printf("the result is %v.", result)
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

func func_name(param1 int, param2 int) int {
	total := 0

	total = param1 + param2
	return total
}
