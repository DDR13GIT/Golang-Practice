package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang!! ")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	fruitList = append(fruitList, "Banana", "Kiwi")
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Limited fruitlist:", fruitList[1:4])

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println("High scores:", highScores[1:])
	

	highScores = append(highScores, 555, 666, 777)
	fmt.Println("High scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from a slice based on index

	var courses = []string{"react", "angular", "vue", "svelte"}
	fmt.Println("Courses:", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing:", courses)
}
