package main

import "fmt"

// first letter of variable if capital means it is public and can be accessed from other packages
const LoginToken string = "asdjfklasdjfklasdjfkl"

func main() {
	var username string = "Debopriya Deb Roy"
	fmt.Println("Hello, " + username + "!")
	fmt.Println("How are you doing?")
	fmt.Printf("This is a variable of type %T \n", username)

	var isLoggedIn bool = true
	
	fmt.Printf("This is a variable of type %T \n", isLoggedIn)

	//default values and aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("This is a variable of type %T \n", anotherVariable)

	//implicit type
	var website = "https://www.debroy.com"
	fmt.Println(website)
	
	//no var style
	age := 25  //this is not allowed outside a function
	fmt.Println(age)

	fmt.Println(LoginToken)
}
