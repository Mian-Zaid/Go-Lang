package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//basic for loop
	fmt.Println("\nBasic For Loop!")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//For loop with Range
	fmt.Println("\nFor Loop with Range!")
	for i := range days { //i has Index, not the Value
		fmt.Println(days[i])
	}

	//For Each loop
	fmt.Println("\nKind of ForEach Loop!")
	for i, day := range days { //i has Index, not the Value
		fmt.Printf("Index is %v and Value is %v\n", i, day)
	}

	fmt.Println("\nKind of While loop!")
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 7 {
			break //break out of loop
		}

		if rougueValue == 6 {
			goto endOfCode
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

endOfCode:
	fmt.Println("Jumped to the End of Code!")
}
