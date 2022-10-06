package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("\n--------------------------------")
	//basic declaration

	const url = "https://lco.dev"

	//create a request
	response, err := http.Get(url) //you can use POST also
	checkNilError(err)

	fmt.Printf("Response type: %T", response)

	defer response.Body.Close()

	//read a response
	dataInBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println("Data read: ", string(dataInBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
