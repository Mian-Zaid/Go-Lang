package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my shop")
	fmt.Println("Please give rating to my services between 1 & 5")

	//take input
	reader := bufio.NewReader(os.Stdin)

	rating, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error while taking input, ", err)
	} else {

		numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Thanks for rating, ", numRating+1)
		}
	}

}
