package main

import (
	"fmt"
)

func print() {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine instance...")
	}
}

func worker(done chan bool){
	   
}

func main() {
	//go print()
	//fmt.Println("main function running...")
	//time.Sleep(3 * time.Second)

	ch := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}

		close(ch)
	}()

	/*
		for n := range ch {
			fmt.Printf("n = %d\n", n)
		}*/

	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	/* Buffered Channels */
	messageStore := make(chan string, 2) /* stores 2 string values */

	messageStore <- "buffered" /* no need to create a go routine to input the values
	in a buffered channel*/
	messageStore <- "channel"

	fmt.Println(<-messageStore)
	fmt.Println(<-messageStore)

	/*channels can be used for sync across go routines
	For this purpose, wait groups can be used
	Basically, waiting for one go routine to finish
	if the other has already finished*/

	



}
