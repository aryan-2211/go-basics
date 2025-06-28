package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "peach"

	fmt.Println("the fruit list is", fruitList)
	fmt.Println("the length of fruit list is", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
}
