package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//Loop using increments
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//Loop just for iterating over range, it still returns index
	for i := range days {
		fmt.Println(days[i])
	}

	//in place of underscore there is the index
	for _, day := range days {
		fmt.Println(day)
	}

	value := 1

	for value < 10 {

		if value == 2 {
			goto label
			//after goto is executed it will simply break the loop 
		}

		if value == 5 {
			break
			// value++
			// continue
		}
		fmt.Println(value)
		value++
	}

	//This is a label created to use in goto
label:
	fmt.Println("This statement is executed through goto statement")

}
