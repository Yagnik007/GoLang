package main

import "fmt"

func main() {
	fmt.Println("If else")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 10&1 == 1 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	//We can initialize and use conditionels at the same time
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than equal to 10")
	}
}
