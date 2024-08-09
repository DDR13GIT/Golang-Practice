package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in golang!! ")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	fruitList = append(fruitList, "Banana", "Kiwi")
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Limited fruitlist:", fruitList[1:4])


}
