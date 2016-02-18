package main

import "testing"

func TestVal(t *testing.T) {
	val := myfunc()
	if val != 8 {
		t.Error("Given Value is not 8")
		t.Fail()
	}

}
