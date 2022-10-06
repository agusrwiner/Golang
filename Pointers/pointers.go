package main

import "fmt"

func main() {
	fmt.Println("\n--------------------------------")
	//basic declaration

	number := 35
	ptr_number := &number //'&varName' returns the storaged location of varName

	fmt.Println("Value: ", number, "\tMemory adress (pointer): ", ptr_number)
	fmt.Println("Value stored in ", ptr_number, ": ", *ptr_number) //'*pointerName' returns the value stored in a pointer
}
