package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circum() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	c := circle{radius: 10}

	fmt.Println("area: ", c.area())
	fmt.Println("perim:", c.circum())

}
