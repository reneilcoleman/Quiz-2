package main

import (
	"fmt"
	"sync"
	"time"
)

func waitgroup() {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Executing Gorputine")
	wg.Done()

}

func main() {
	fmt.Println("Go Waitgroup Tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go waitgroup(&wg)
	wg.Wait()
	fmt.Println("Finished Executing my go program")
}
