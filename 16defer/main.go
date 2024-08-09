package main

import "fmt"

func main(){
	defer fmt.Println("This will be printed at the end! Because it is deferred!")
	fmt.Println("Welcome to defer in golang!!")
	fmt.Println("This will be printed first!!")
	fmt.Println("This will be printed second!!")
}