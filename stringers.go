package main

import "fmt"

type database struct {
	name  string
	age   int
	place string
}

func (d database) String() string {
	return fmt.Sprintf("I'm %v and i'm %v years old. I come from %v.", d.name, d.age, d.place)
}
func main() {
	d := database{"Vishwas", 20, "Udupi"}
	fmt.Println(d)
}
