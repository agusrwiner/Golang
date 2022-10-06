package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("\n--------------------------------")
	//basic declaration

	content := "We want to write this inside a file"

	//creation
	file, err := os.Create("./myFile.txt") //'./' means that you will create the file in the same directory

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	//writing
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("The length is: ", length)
	defer file.Close()

	//reading file
	readFile(file.Name())
}

func readFile(fileName string) {
	//when you read a file its read in bytes
	dataInBytes, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("Text data (in Bytes) is: \n\t", dataInBytes)
	fmt.Println("Text data is: \n\t", string(dataInBytes))
}

// this is a common practise
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
