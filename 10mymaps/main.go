package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	//Deleting in maps or slice or other DS happens like this
	delete(languages, "JS")
	fmt.Println(languages)

	//Iterating over map, we can use comma ok error syntax here.
	for _, value := range languages{
		fmt.Println(value)
	}
}
