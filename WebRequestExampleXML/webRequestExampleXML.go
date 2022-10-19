package main

import (
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

	var breakfast_menu1 BreakfastMenu                     //create the reference to a struct
	err := xml.Unmarshal(bytesResponse, &breakfast_menu1) //we save the bytes into the struct
	checkNilError(err)

	fmt.Printf("\nbreakfast_menu1 type: '%T'\n", breakfast_menu1)
	fmt.Println("\nbreakfast_menu1: \n", breakfast_menu1)

	fmt.Printf("\nbreakfast_menu1.FoodList type: '%T'\n", breakfast_menu1.FoodList)
	fmt.Println("\nbreakfast_menu1.FoodList: \n", breakfast_menu1.FoodList)

	fmt.Println("\n========================================================")
	breakfast_menu1.showData()
}

type BreakfastMenu struct {
	//XMLName  xml.Name `xml:"breakfast_menu"`
	FoodList []Food `xml:"food"`
}

func (bm *BreakfastMenu) showData() {
	fmt.Println("\n--------------------------------")
	fmt.Println("Food list: ")
	for _, food := range bm.FoodList {
		food.showData()
	}
	fmt.Println("--------------------------------")
}

type Food struct {
	Name        string `xml:"name"`
	Price       string `xml:"price"`
	Description string `xml:"description"`
	Calories    string `xml:"calories"`
}

func (f *Food) showData() {
	//fmt.Println("\n--------------------------------")
	fmt.Println("\tName: ", f.Name)
	fmt.Println("\tPrice: ", f.Price)
	fmt.Println("\tCalories: ", f.Calories)
	fmt.Println("\tDescription: ", f.Description)
	fmt.Println("\t--------------------------------")
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
