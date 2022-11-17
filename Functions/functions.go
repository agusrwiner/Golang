package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration
	//func functionName(arguments){}

	fmt.Println("--------------------------------")
	greeter()

	fmt.Println("--------------------------------")
	result1 := adder(10, 15)
	fmt.Println("The result of the sum is: ", result1)

	fmt.Println("--------------------------------")
	result2 := proAdder(10, 15, 23, 25, 100)
	fmt.Println("The result of the PRO sum is: ", result2)
}

//========== basic function ==========
func greeter() {
	fmt.Println("Hellooooo")
}

//========== with parameters ==========
//func funcName(param1 param1Type,paramX paramX Type) retornType{}
func adder(n1 int, n2 int) int {
	return n1 + n2
}

//========== with infinite parameters ==========
//func funcName(paramSliceName ...paramTypes) retornType{}
func proAdder(values ...int) int {
	total := 0

	//values its a slice
	for _, value := range values {
		total += value
	}

	return total
}
