package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang!!")
	greet()

	result := add(3, 5)
	fmt.Println("Result is:", result)

	proResult := proAdd(1, 2, 3, 4, 5)
	fmt.Println("Pro result is:", proResult)
}

func greet() {
	fmt.Println("Hi there!! What's up?")
}

func add(x int, y int) int{
	return x + y
}

func proAdd(values ...int) int{  //takes a lot of values of type int
	result := 0
	for _, value := range values{
		result += value
	}
	return result
}