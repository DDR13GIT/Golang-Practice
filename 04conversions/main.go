package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcom to our pizza app!!")
	fmt.Println("Please rate our pizza between 1 and 5:")

	reading := bufio.NewReader(os.Stdin)
	input, _ := reading.ReadString('\n')
	fmt.Println("Thanks for rating our pizza:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
