package main

import "fmt"

//Declaring a public variable:

//In Go, a public variable is typically denoted by starting its name with an uppercase letter. However, it's important to understand that Go doesn't have the concept of "public" or "private" variables in the same way that some other languages like Java or Python do. In Go, whether a variable is accessible outside its package (similar to public) or only within its package (similar to private) depends on its initial letter case: Variables starting with an uppercase letter are exported from the package and accessible from outside the package.Variables starting with a lowercase letter are not exported and are only accessible within the package they are defined in.

const FirstName string = "Yagnik"

var LastName string = "Talaviya"

func main() {
	//entry point of the execution.
	var username string = "yagnik"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal byte = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.243513453534523
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var initialisedVariable int
	fmt.Println(initialisedVariable) //Output: 0
	fmt.Printf("Variable is of type: %T \n", initialisedVariable)

	// implicit type variable declaration: Once this way a variable is declared, the lexer identifies the data type by itself and one cannot change the value of the variable of to a different data type. We still can change the value but of the same data type.

	var website = "fckherzaman.pk"
	fmt.Println(website)

	//no var style: The background process is similar to the implicit type, but the thing is we could define this way under a method and not globally.

	numberOfUser := 1000
	fmt.Println(numberOfUser)

	fmt.Printf("My full name is %s %s", FirstName, LastName)
}
