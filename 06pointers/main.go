package main

import "fmt"

func main() {
	myNumber := 23

	var ptr = &myNumber

	//ptr is the memory address in which myNumber is stored
	//*ptr denotes the content stored in ptr
	fmt.Println("the value of ptr...", ptr)
	fmt.Println("the value of *ptr(the content inside the ptr)", *ptr)

	//the changes made in *ptr will change the actual value of myNumber
	*ptr = *ptr + 2
	fmt.Println("new value...", myNumber)
}
