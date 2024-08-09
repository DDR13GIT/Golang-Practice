package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://reqres.in/api/users?page=2"

func main() {
	fmt.Println("Welcome to urls in golang!!")
	fmt.Println("URL is:", myurl)

	result, _ := url.Parse(myurl)
	fmt.Println("Scheme is:", result.Scheme)
	fmt.Println("Host is:", result.Host)
	fmt.Println("Path is:", result.Path)
	fmt.Println("Query is:", result.RawQuery)

	qparams := result.Query()
	fmt.Println("Query params are:", qparams)

}
