package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang!")
	hello()

	adderRes := adder(1, 3)
	fmt.Println("Adder res is: ", adderRes)

	proRes, mess := proAdder(1, 3, 3, 4)
	fmt.Println("Pro Adder res is: ", proRes)
	fmt.Println("Pro Adder message is: ", mess)

}

func hello() {
	fmt.Println("Hello")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(vals ...int) (int, string) {
	total := 0

	for _, value := range vals {
		total += value
	}

	return total, "Hello from proAdder\n"
}
