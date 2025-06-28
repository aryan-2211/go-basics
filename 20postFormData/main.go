package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("sending form data to server in golang")
	PostFormData()
}

func PostFormData() {
	const myurl = "http://localhost:8000/postform"

	//making the form data
	data := url.Values{}
	data.Add("FirstName", "Aryan")
	data.Add("LastName", "Shandilya")
	data.Add("email", "aryan@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
