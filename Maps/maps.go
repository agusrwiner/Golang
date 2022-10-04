package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//varName := make(map[keyType]valueType)

	personList := make(map[int]string) // a map with intType keys and stringTypes values
	fmt.Println(personList)

	fmt.Println("--------------------------------")
	//========== adding items ==========
	//mapName[key] = value

	personList[1] = "Juan"
	personList[2] = "Pedro"
	personList[3] = "Jose"
	personList[4] = "Ariana"
	fmt.Println("Persons list: ", personList)

	fmt.Println("--------------------------------")
	//========== deleting items ==========
	//delete( mapName,key )

	delete(personList, 2)
	fmt.Println("Persons list (position '2'): ", personList)

	fmt.Println("--------------------------------")
	//========== selecting specific items ==========
	//mapName[key]

	fmt.Println("Persons list (position '3'): ", personList[3])

	fmt.Println("--------------------------------")
	//========== iterating throught the map ==========
	//for key, value := range mapName{}

	for key, value := range personList {
		fmt.Print("\tKey: ", key)
		fmt.Println("\tValue: ", value)
	}
	fmt.Println("Persons list: ", personList)
}
