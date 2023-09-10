package main

import "fmt"

func main() {
	fmt.Println("Welcome to GO Structs!")

	//no inheretance in golang; No super or parent

	zaid := User{"Muhammad Zaid Ali", "Zaid@go.dev", true, 25}
	fmt.Println(zaid)
	fmt.Printf("\nZaid details are %+v\n", zaid)
	fmt.Printf("\nName is %v and Email is %v\n", zaid.Name, zaid.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
