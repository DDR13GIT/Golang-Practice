package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang!!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days of the week:", days)

	for i:=0; i<len(days); i++{
		fmt.Println("Day:", days[i])
	}

	for i := range days{
		fmt.Println("Day:", days[i])
	}

	for index, value := range days{
		fmt.Printf("Index:%v Day:%v\n", index, value)
	}

	for _, value := range days{
		fmt.Printf("Day:%v\n", value)
	}

	rougeValue := 1

	for rougeValue < 10{
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

}
