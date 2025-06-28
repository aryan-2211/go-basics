package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}

	fruitList = append(fruitList, "apple", "mango", "banana")
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 45
	highScores[1] = 25
	highScores[2] = 65
	highScores[3] = 15

	highScores = append(highScores, 32, 87, 94)
	sort.Ints(highScores)

	fmt.Println(highScores)
}
