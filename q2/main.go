package main

import (
	"fmt"
	"time"
)

func myhobbies(s chan string) {
	fmt.Println("Starting program")
	time.Sleep(1 * time.Second)
	s <- "1. Basketball"
	fmt.Println("Program finished")

}

func main() {
	fmt.Println("My everyday hobbies are: ")

	channel := make(chan string, 2) // add to to allow the buffered channel to run. since we only had one receive statement, adding two will will now allow the second go routine to run.
	defer close(channel)            // This will stop any go routines from sending or receiving to the channel when its closed.

	go myhobbies(channel)
	go myhobbies(channel) // we now have two go routines trying to send to the same channel.

	val := <-channel
	fmt.Println(val)

	time.Sleep(1 * time.Second)
}
