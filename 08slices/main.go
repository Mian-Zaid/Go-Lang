package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to GO Slices!")

	var fruitList = []string{"Apple", "Mango", "Apricot"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Tomato", "Banana")
	fmt.Println("fruit list is ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("fruit list is ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 1
	highScores[1] = 5
	highScores[2] = 3
	highScores[3] = 4
	// highScores[4] = 7

	highScores = append(highScores, 8, 10, 66)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println("sorted scores: ", highScores)

	//remove element from slice using index
	var courses = []string{"Python", "C++", "JavaScript", "Java", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
