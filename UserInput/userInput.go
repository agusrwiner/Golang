package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Whats your name? >")

	//varName,error := reader.ReadString('enderCharacter')
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) //you need to take out the \n of the input

	fmt.Println("Hi ", input, " welcome!")

	fmt.Println("--------------------------------")
	//========== data conversion ==========

	fmt.Println("Enter a numer: ")
	aux1, _ := reader.ReadString('\n')
	aux1 = strings.TrimSpace(aux1)
	num1, err := strconv.ParseFloat(aux1, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number you entered plus 10 is: ", num1+10)
	}
}
