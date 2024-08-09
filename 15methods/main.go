package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang!! ")

	//no inheritance in go
	//no super or parent keywords

	user1 := User{"Debopriya", "Deb Roy", "ddroy13@gmail.com", true, 25}
	fmt.Println("User1:", user1)
	fmt.Printf("User1:%+v\n", user1)

	user1.GetStatus()
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
	Age       int
}

func (u User) GetStatus() {
	fmt.Println(u.FirstName, "Is Active:", u.IsActive)
}
