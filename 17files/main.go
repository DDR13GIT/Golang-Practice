package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang!!")

	content := "This needs to go in a file"

	file, err := os.Create("./mysample.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mysample.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("File content is: ", string(data))
}

func checkNilErr(err error){
	if err != nil {
		fmt.Println("Error reading file!!")
		return
	}
}