package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
<feed>
	<item>
		<title>postTitle</title>
		<link>postLink</link>
		<description>postDescription</description>
		<pubDate>Mon, 17 Oct 2022 14:13:37 +0000</pubDate>
		<dc:creator>CISA</dc:creator>
		<guid isPermaLink="false">18076 at https://us-cert.cisa.gov</guid>
	</item>
</feed>
*/

func main() {
	fmt.Println("\n########################################################")
	fmt.Println("########################################################")
	fmt.Println("########################################################")

	const url = "https://www.cisa.gov/uscert/ncas/bulletins.xml"
	feed1, err := Fetch(url)
	checkNilError(err)

	//====== TESTS ======
	fmt.Printf("\tfeed type: '%T'\n", feed1)
	fmt.Println("\tfeed: \n", feed1)

	fmt.Printf("\tfeed.FoodList type: '%T'\n", feed1.Posts)
	fmt.Println("\tfeed.FoodList: \n", feed1.Posts)

	fmt.Println("\n========================================================")
	feed1.showData()
}

// ========= RSS =========
type RSS struct {
	Channel string `xml:"rss"`
}

// ========= FEED =========
type Feed struct {
	//Name  string //could also be an id
	XMLName xml.Name `xml:"channel"`
	Posts   []Post   `xml:"item"`
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

// ========= POST =========
type Post struct {
	//Id          int
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	//Description string `xml:"description"`
}

func (p *Post) showData() {
	//fmt.Println("\n--------------------------------")
	//fmt.Println("\tId: ", p.Id)
	fmt.Println("\tTitle: ", p.Title)
	//fmt.Println("\tDescription: ", p.Description)
	fmt.Println("\t--------------------------------")
}

// ========= FUNCS & AIDS =========
func Fetch(url string) (Feed, error) {
	response, err := http.Get(url) //we make the request
	checkNilError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body) //we get an slice of bytes
	checkNilError(err)

	var feed Feed
	err = xml.Unmarshal(bytes, &feed) //we parse the slice of bytes to a Feed struct

	return feed, err //we return the Feed
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
