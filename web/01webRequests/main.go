package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	url = "https://lco.dev"
)

func main() {
	fmt.Println("Welcome to web request in GoLang!")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of respone is : %T\n", response)

	defer response.Body.Close() //caller's responsibilty to close it

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("Response is :", content)
}
