package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("the type of response is...%T", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
