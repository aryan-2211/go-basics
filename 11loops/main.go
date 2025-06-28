package main

import "fmt"

func main() {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thurs", "Fri", "Sat"}

	//type 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//type 2
	for i := range days {
		fmt.Println(days[i])
	}

	//type 3
	for index, day := range days {
		fmt.Printf("the index is %v and the day is %v.", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		fmt.Println("value is...", rougueValue)
		rougueValue++
	}

}
