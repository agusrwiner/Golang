package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("\n--------------------------------")
	//basic declaration

	const url1 = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gsdg345"

	//parse information from a URL
	result, err := url.Parse(url1)
	checkNilError(err)
	fmt.Println("Full Result: ", result)
	fmt.Println("Result Host: ", result.Host)
	fmt.Println("Result Path: ", result.Path)
	fmt.Println("Result RawQuery: ", result.RawQuery)

	queryParameters := result.Query()
	fmt.Printf("\nQuery Params type: %T\n", queryParameters)
	fmt.Println("Query Params (its a Map): ", queryParameters)
	fmt.Println("")

	//using a for loop with the response
	for key, value := range queryParameters {
		fmt.Println("Key: ", key, "\tValue: ", value)
	}

	fmt.Println("\n--------------------------------")
	//constructing a URL

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	url2 := partsOfUrl.String()
	fmt.Println("URL constructed: ", url2)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
