package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	//basic declaration

	currentTime := time.Now()
	fmt.Println("The current time is: ", currentTime)

	//========== format ==========
	//see the documentation to see other formats

	formatedTime := currentTime.Format("02/01/2006 15:04:05 Monday")
	fmt.Println("The current (formated) time is: ", formatedTime)

	fmt.Println("--------------------------------")
	//========== create dates ==========

	myBirthday := time.Date(2004, time.June, 06, 0, 0, 0, 0, time.UTC)
	fmt.Println("My birthday is: ", myBirthday.Format("02/01/2006 15:04:05 Monday"))
}
