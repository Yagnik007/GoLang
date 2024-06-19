package main

import "fmt"

func main() {
	fmt.Println("These are functions")
	result := adder(10, 90)
	fmt.Println(result)

	//To take multiple return arguments we simply use comma ok syntax, if we don't care about any of the result we can sumply replace it with underscore
	result, myMessage := proAdder(1, 2, 3, 12, 3, 4, 5, 65, 123, 3, 5, 6, 6)
	fmt.Println(result)
	fmt.Println(myMessage)
}

//A proper way to define a function, it could  be defined anywhere
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//Variadic functions can take n number of arguments and can also return multiple results
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi, proAdder function here!!"
}
