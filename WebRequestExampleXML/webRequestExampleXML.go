package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
<breakfast_menu>
	<food>
		<name>Belgian Waffles</name>
		<price>$5.95</price>
		<description>Two of our famous Belgian Waffles with plenty of real maple syrup</description>
		<calories>650</calories>
	</food>
</breakfast_menu>
*/

func main() {
	fmt.Println("\n########################################################")
	fmt.Println("########################################################")
	fmt.Println("########################################################")

	const url = "https://www.w3schools.com/xml/simple.xml"

	//create a request
	bytesResponse := getContentInBytes(url)

	var breakfast_menu1 Breakfast_menu                    //create the reference to a struct
	err := xml.Unmarshal(bytesResponse, &breakfast_menu1) //we save the bytes into the struct
	checkNilError(err)

	jsonData, err := json.Marshal(breakfast_menu1) //parse our data to json
	checkNilError(err)

	//fmt.Println("JSON data:\n\t", jsonData)
	fmt.Println(" ")
	fmt.Println(string(jsonData))
}

type Breakfast_menu struct {
	//XMLName  xml.Name `xml:"breakfast_menu"`
	FoodList []Food `xml:"food"`
}
type Food struct {
	Name        string `xml:"name"`
	Price       string `xml:"price"`
	Description string `xml:"description"`
	Calories    string `xml:"calories"`
}

// func getContentAsString(url string) string {
// 	response, err := http.Get(url) //you can use POST also
// 	checkNilError(err)
// 	defer response.Body.Close()

// 	bytes, err := ioutil.ReadAll(response.Body) //slice of bytes
// 	checkNilError(err)

//		return string(bytes)
//	}
func getContentInBytes(url string) []byte {
	response, err := http.Get(url) //you can use POST also
	checkNilError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body) //slice of bytes
	checkNilError(err)

	return bytes
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
