package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("--------------------------------")
	/*type structName struct{
		varName varType
		varName varType
	}*/

	user1 := User{"Agustin", "agus@gmail.com", 18}
	fmt.Println(user1)
	fmt.Println("Name: ", user1.Name, "\n")

	user1.showUserData()
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) showUserData() {
	fmt.Println("Name: ", u.Name)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Age: ", u.Age)
	fmt.Println("--------------------------------")
}
