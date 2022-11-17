package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration
	//var varName [maxCapacity]dataType
	var fruitList [4]string
	fruitList[0] = "Banana"
	fruitList[1] = "Peach"
	fruitList[2] = "Apple"
	fmt.Println("Fruit list: ", fruitList, " List lenght: ", len(fruitList))

	var vegetablesList = [3]string{"Lettuce", "Potato", "Carrot"}
	fmt.Println("Fruit list: ", vegetablesList, " List lenght: ", len(vegetablesList))
	fmt.Println("")
}
