package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of our Pizza: ")

	// comma ok | err (Used instead of try-catch(Not available), because of the design of the language the errors are treated as merely another variables or like that)

	input, _ := reader.ReadString('\n') //This line right here represents (ok , err) where if input is error-free, it will be stored in input variable and if the error occurs it will be stored in the variable next to comma. In many cases we do not care about the error and hence "underscore" is a replacment for that. In some cases we might care only about error and we could similarly replace input variable with underscore and store error in a variable.

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input) //String

}
