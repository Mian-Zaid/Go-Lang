package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else in Golang!")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println("Result is: ", result)

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less then 10")
	}

}
