package main

import	"fmt"
func add(a int,b int) int {
	return a+b;
}

func main() {
	var a,b int;
	a=10;
	b=90;
	fmt.Println("Sum of A and B is: " , add(a,b))
}
