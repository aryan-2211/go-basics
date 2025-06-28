package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("How to create JSON data in golang")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "YT.com", "abc123", []string{"webdev", "JS"}},
		{"Mern", 99, "LT.com", "bcd123", []string{"fullstack", "JS"}},
		{"Angular", 299, "YT.com", "xyz123", nil},
	}

	//package this data as json data
	//json.Marshal is the default that can be used
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

//decoding json data later
