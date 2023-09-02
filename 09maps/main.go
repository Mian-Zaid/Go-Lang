package main

import "fmt"

func main() {
	fmt.Println("Welcome to GO MAPS!")

	//make map
	languages := make(map[string]string)

	//add data to map
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	//delete element from map
	delete(languages, "RB")
	fmt.Println("List of all Languages: ", languages)

	//loop through the map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
