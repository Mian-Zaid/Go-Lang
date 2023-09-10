package main

import "fmt"

func main() {
	fmt.Println("Welcome to the GO Pointers!")

	// var ptr *int
	// fmt.Println("Value of ptr is ", ptr)

	myNumber := 23
	var numPtr = &myNumber
	fmt.Println("Value of actual ptr is ", numPtr)
	fmt.Println("Value of actual ptr is ", *numPtr)

	*numPtr = *numPtr + 2
	fmt.Println("new value of actual ptr is ", myNumber)

}
