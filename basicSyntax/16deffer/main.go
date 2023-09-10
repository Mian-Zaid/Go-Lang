package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("World")

	fmt.Println("Welcome to deffer in Golang!")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
