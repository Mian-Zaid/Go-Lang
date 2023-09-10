package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in GOLANG")

	currenttime := time.Now()
	fmt.Println("Current time: ", currenttime)
	fmt.Println("Formated time: ", currenttime.Format("01-02-2006 Monday"))

}
