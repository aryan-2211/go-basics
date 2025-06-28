package main

import "fmt"

func main() {
	var username string = "name of the user"
	fmt.Println(username)
	fmt.Printf("the type of the variable is: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of the variable is: %T \n", isLoggedIn)
}
