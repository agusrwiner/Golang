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
	breakfasMenu1, err := Fetch(url)
	checkNilError(err)

	json, err := Serialize(breakfasMenu1)
	checkNilError(err)

	fmt.Printf("\nbreakfast_menu1 type: '%T'\n", breakfasMenu1)
	fmt.Println("\nbreakfast_menu1: \n", breakfasMenu1)

	fmt.Printf("\njson type: '%T'\n", json)
	fmt.Println("\njson: \n", string(json))

	fmt.Println("\n========================================================")
	breakfasMenu1.showData()
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

/*
Recieves URL of an XML -
Retorns the XML parsed to an BreakfastMenu
*/
func Fetch(url string) (BreakfastMenu, error) {
	response, err := http.Get(url) //we make the request
	checkNilError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body) //we get an slice of bytes
	checkNilError(err)

	var bm BreakfastMenu
	err = xml.Unmarshal(bytes, &bm) //we parse the slice of bytes to a Feed struct

	return bm, err //we return the Feed
}

// serializar a json
func Serialize(breakfastMenu BreakfastMenu) ([]byte, error) {
	jsonInBytes, err := json.Marshal(breakfastMenu)

	return jsonInBytes, err
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
