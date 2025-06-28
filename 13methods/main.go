package main

import "fmt"

func main() {
	abc := User{"abc", "abc@go.dev", true, 21}

	fmt.Println("the email is...", abc.Email)
	abc.GetStatus() 
	abc.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active..", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@godev"
	fmt.Println("the new email is..", u.Email)
}
