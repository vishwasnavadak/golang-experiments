package main

import (
	"fmt"
	"testing"
)

func sampleTestCase(t *testing.T) {
	tVal := []int{1, 2, 3, 4, 5, 6}
	for _, val := range tVal {
		if val != 8 {
			t.Error("8 is not found in the given array")
		}
	}
}
func main() {
	fmt.Println("Hello World")
}
