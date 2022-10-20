package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(5, 4)
	if sum != 9 {
		t.Error("Funtion return incorrect")
		t.Fail()
	} else {
		t.Log("Test passed")
	}
}

func Add(n1, n2 int) int {
	//return n1 + n2
	return 45
}
