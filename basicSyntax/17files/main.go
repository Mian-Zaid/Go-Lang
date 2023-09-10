package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!")

	fileName := "myFile.txt"

	createFile(fileName)

	readFile(fileName)
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	checkNilErr(err)

	content := "This is written in file only - MZA.com"
	length, err := file.WriteString(content)
	checkNilErr(err)

	fmt.Println("The lenght written is: ", length)
	defer file.Close()
}

func readFile(fileName string) {
	dataBytes, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Data in file is \n", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
