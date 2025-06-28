package main

import "fmt"

func main() {
	//structs are the alternate for classes in golang
	//there is no inheritance in go
	//no super or parent class

	aryan := User{"Aryan", "aryan@go.dev", 21, true}
	fmt.Println(aryan)

	honda := car{"city", "black", 12345, "25/06/25", true}
	fmt.Println("the car's name is", honda.Name)

	fmt.Println("the name is..", aryan.Name)
	fmt.Println("the age is..", aryan.Age)
}

type car struct {
	Name     string
	Colour   string
	Number   int
	Date     string
	Validity bool
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
