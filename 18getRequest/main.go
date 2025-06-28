package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myurl = "http://localhost:8000/get"

func main() {
	fmt.Println("How to make a get request")
	makeGetRequest()
}

func makeGetRequest() {

	//get the response
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	//read the response
	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)

}
