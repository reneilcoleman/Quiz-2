package main

import (
	"fmt"
	"time"
)

func count(c string) {
	for i := 0; i < 2; i++ {
		fmt.Println(c)
		time.Sleep(time.Microsecond)

	}
}

func main() {
	go count("30 Teams are in the NBA")
	count("There are 15 in the west and 15 in the east")

	time.Sleep(time.Second)
}
