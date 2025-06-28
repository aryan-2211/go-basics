package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this needs to be inserted in the file"

	file, err := os.Create("./demoFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println(length)
	defer file.Close()

	readFile("./demoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
