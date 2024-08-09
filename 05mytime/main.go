package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to golang study time!")
	presentTime := time.Now()
	fmt.Println("The current time is: " + presentTime.String())
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2024, time.March, 13, 11, 40, 20, 215, time.Local)
	fmt.Println("Created date is: " + createDate.String())
	fmt.Println("Created date is: " + createDate.Format("01-02-2006 15:04:05 Monday"))
}

// GOOS="windows" go build
// GOOS="linux" go build
// GOOS="darwin" go build for mac

