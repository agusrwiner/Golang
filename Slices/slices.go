package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration
	//var varName = []varType{}
	//(you can declare it empty) EXAMPLE:
	//var fruitList = []string{}

	var fruitList = []string{"Apple", "Orange", "Banana"}
	fmt.Println("Fruit list: ", fruitList, " List lenght: ", len(fruitList))

	fmt.Println("--------------------------------")
	//========== adding items to the slice ==========
	//sliceName = append( sliceName,dataToAdd )
	fruitList = append(fruitList, "Pear", "Mango")
	fmt.Println("Fruit list: ", fruitList)

	fmt.Println("--------------------------------")
	//========== DELETE items from the slice based on INDEX ==========
	//you use de 'append' function to retrieve only the desired part
	//you use the '...' as a default

	index := 2
	fruitList = append(fruitList[:index], fruitList[index+1:]...)

	fmt.Println("Fruit list (with an item removed): ", fruitList)

	fmt.Println("--------------------------------")
	//========== sub slices ==========
	//you can select pieces of an slice

	//                        from a position to a position
	var auxList = append(fruitList[1:3])
	fmt.Println("Fruit list (cropped): ", auxList)

	fmt.Println("--------------------------------")
	//========== make function ==========
	//listName := make( []varType,max capacity )

	numbers := make([]int, 4)
	numbers[0] = 100
	numbers[1] = 146
	numbers[2] = 1064
	numbers[3] = 1001

	fmt.Println("Numers list: ", numbers)

	numbers = append(numbers, 587, 247, 35)
	fmt.Println("Numers list: ", numbers)

	fmt.Println("--------------------------------")
	//========== functions ==========

	//sort.Ints
	sort.Ints(numbers)
	fmt.Println("Numers list (sorted): ", numbers)

	fmt.Println("")
}
