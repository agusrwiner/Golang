package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
<feed>
	<post>
		<title>postTitle</title>
		<description>postDescription</description>
	</post>
</feed>
*/

func main() {
	fmt.Println("\n########################################################")
	fmt.Println("########################################################")
	fmt.Println("########################################################")

	/*rawXmlData :=
	`<feed>
		<post>
			<title>Virus H</title>
			<description>Este virus afecta a los usuarios de Postgre DB</description>
		</post>
		<post>
			<title>Virus W-YZ</title>
			<description>Este virus afecta a los usuarios de Windows 8 en adelante</description>
		</post>
		<post>
			<title>Bug R2-D</title>
			<description>Este bug afecta a los usuarios fans de SW</description>
		</post>
		<post>
			<title>Virus XUW32</title>
			<description>Este virus afecta a los usuarios de Linux</description>
		</post>
	</feed>`*/

	const url = ""
	bytesResponse := getContentInBytes(url)

	var feed1 Feed
	err := xml.Unmarshal(bytesResponse, &feed1)
	checkNilError(err)

	//====== TESTS ======
	fmt.Printf("\tfeed type: '%T'\n", feed1)
	fmt.Println("\tfeed: \n", feed1)

	fmt.Printf("\tfeed.FoodList type: '%T'\n", feed1.Posts)
	fmt.Println("\tfeed.FoodList: \n", feed1.Posts)

	fmt.Println("\n========================================================")
	feed1.showData()
}

// ========= POST =========
type Post struct {
	//Id          int
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

func (p *Post) showData() {
	//fmt.Println("\n--------------------------------")
	//fmt.Println("\tId: ", p.Id)
	fmt.Println("\tTitle: ", p.Title)
	fmt.Println("\tDescription: ", p.Description)
	fmt.Println("\t--------------------------------")
}

// ========= FEED =========
type Feed struct {
	//Name  string //could also be an id
	Posts []Post `xml:"feed"`
}

func (feed *Feed) showData() {
	fmt.Println("\n--------------------------------")
	//fmt.Println("Name: ", feed.Name)
	fmt.Println("Posts list: ")
	for _, post := range feed.Posts {
		post.showData()
	}
	fmt.Println("--------------------------------")
}

// ========= FUNCS & AIDS =========
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
