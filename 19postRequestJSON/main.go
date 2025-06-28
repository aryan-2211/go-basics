package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Performing POST request in golang")
	PerformPostRequestJSON()
}

func PerformPostRequestJSON() {
	const myurl = "http://localhost:8000/post"

	//making fake JSON payload using the strings package & NewReader fn
	requestBody := strings.NewReader(`
		{
			"courseName" : "learning go",
			"price" : "0",
			"platform" : "Youtube"
		}
	`)

	//important
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
