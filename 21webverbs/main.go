package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs in golang!!")
	//PerformGetRequest()
	PerformPostJSONRequest()
}

func PerformGetRequest() {
	const myurl string = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("GET Response Status:", response.Status)
	fmt.Println("Content length", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body) //content is a byte slice
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is ", byteCount)
	fmt.Println("Content:", responseString.String())
	//fmt.Println("Content:", string(content))

}

func PerformPostJSONRequest() {
	const myurl string = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
	"coursename":"Let's go with golang",
	"price":10,
	"platform":"Udemy"
	}`)

	response, _ := http.Post(myurl, "application/json", requestBody)

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content is:", string(content))

}
