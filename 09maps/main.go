package main

import "fmt"

/*
func main() {
	//syntax
	languages := make(map[string]string)

	//entering values in the map(key, value)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println(languages)

	//extracting value
	fmt.Println(languages["RB"])

	//deleting
	delete(languages, "PY")
	fmt.Println(languages)
}*/

func main() {
	Mymap := make(map[int]string)

	Mymap[1] = "abcd"
	Mymap[2] = "def"
	Mymap[3] = "xyz"

	value, ok := Mymap[4]
	if ok {
		fmt.Println("value exists in map", value)
	} else {
		fmt.Println("value doesn't exist in map")
	}

	delete(Mymap, 1)
	fmt.Println(Mymap)
}
