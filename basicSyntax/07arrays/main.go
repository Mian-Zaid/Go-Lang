package main

import "fmt"

func main() {
	fmt.Println("Welcome to GO Arrays!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Tomato"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list len is ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroo"}
	fmt.Println("Veg List is ", vegList)
	fmt.Println("Veg List len is ", len(vegList))
}
