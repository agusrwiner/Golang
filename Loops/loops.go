package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)

	fmt.Println("")
	//basic for
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	println("")

	fmt.Println("")
	//range iterator
	for j := range days {
		fmt.Println(" ", j, " - ", days[j])
	}

	fmt.Println("")
	//foor each (kinda)
	for index, value := range days {
		fmt.Print("Index: ", index)
		fmt.Println("\tValue: ", value)
	}

	fmt.Println("")
	//while (you use de for)

	num := 0
	for num < 10 {
		if num == 3 || num == 6 {
			num++
			continue
		} else if num == 8 {
			break
		}

		fmt.Print(num, " ")
		num++
	}
	fmt.Println("")
}
