package main

import (
	"fmt"
)

type Findarea interface {
	area() float64
}

type Triangle struct {
	base, height float64
}

//Triangle Method
func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func print(p Findarea) {
	fmt.Println(p)
	fmt.Println(p.area())

}

func main() {
	t := Triangle{base: 10, height: 15}

	print(t)

}
