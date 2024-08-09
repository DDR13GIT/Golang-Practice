package main

import "fmt"

func main() {
	fmt.Println("Welcome to mymaps in golang!! ")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("Languages:", languages)
	fmt.Println("Go is short for", languages["GO"])

	delete(languages, "RB")
	fmt.Println("Languages:", languages)	

	for key, value := range languages{
		fmt.Println("Key is:", key, "Value is:", value)
	}
}
