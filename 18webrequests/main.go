package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Welcome to web requests in golang!!")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Response is of type: %T\n", response)


	defer response.Body.Close() //closing the response body is the responsibility of the caller

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Content of the website is: ", string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error reading file!!")
		return
	}
}
