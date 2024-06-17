package main

import "fmt"

func main() {
	fmt.Println("These are pointers")

	// This is a pointer that points to the value of type int being stored at the address
	var ptr *int

	fmt.Println("Value of pointer is ", ptr) //nil if the value is not assigned to the pointer

	myNumber := 40

	var ptrMyNumber = &myNumber

	fmt.Println("Value of the pointer of a number is ", ptrMyNumber) //Returns the address
	fmt.Println("Value of the pointer of a number is ", *ptrMyNumber) //Returns the value present at the address

	*ptrMyNumber += 2 //Operation performed on the values, not on the copies of value.
	fmt.Println("New value is ", myNumber)
}
