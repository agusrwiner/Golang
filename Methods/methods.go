package main

import "fmt"

func main() {
	user1 := User{"Agustin", "agus@gmail.com", 18}
	user2 := User{"Martina", "marti@gmail.com", 40}
	user3 := User{"Pedro", "pedro@gmail.com", 15}

	fmt.Println("")
	fmt.Println("--------------------------------")
	user1.showUserData()
	user2.showUserData()
	user3.showUserData()

	user1.isOldEnough()
	user2.isOldEnough()
	user3.isOldEnough()
	fmt.Println("")
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetAge() {
	fmt.Println("User Status: ", u.Age)
}

//We can see that this function recieves a copy of the Struct,
//so it doesnt actually change the Email of the user, for that
//we would need to pass a Pointer to the Struct
func (u User) NewEmail() {
	u.Email = "newEmail@gmail.com"
	fmt.Println("User Email: ", u.Email)
}
func (u User) showUserData() {
	fmt.Println("Name: ", u.Name)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Age: ", u.Age)
	fmt.Println("--------------------------------")
}
func (u User) isOldEnough() {
	if u.Age >= 18 {
		fmt.Println("Yes, the user is ", u.Age, " years old")
	} else {
		fmt.Println("No, the user is ", u.Age, " years old")
	}
}
