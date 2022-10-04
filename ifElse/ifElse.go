package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration

	temperature := 2
	messagge := ""

	if temperature >= 20 && temperature < 30 {
		messagge = "Its nice"
	} else if temperature > 30 {
		messagge = "Its smoking hot"
	} else {
		messagge = "Its cold"
	}
	fmt.Println(messagge)
	fmt.Println("Its ", temperature, " degrees.")

	fmt.Println("--------------------------------")
	//========== switch ==========

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5) + 1
	fmt.Println("Random number: ", randomNumber, "\n ")

	switch randomNumber {
	case 1:
		fmt.Println("The number is: 1")
		fallthrough
	case 2:
		fmt.Println("The number is: 2")
		fallthrough
	case 3:
		fmt.Println("The number is: 3")
		fallthrough
	case 4:
		fmt.Println("The number is: 4")
		fallthrough
	case 5:
		fmt.Println("The number is: 5")
		fallthrough
	default:
		fmt.Println("This is de default option")
	}
}
