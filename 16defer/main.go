package main

import "fmt"

//Defer keyword pushes that line into a stack where the stack will execute each deferred line at the last of the function just before returning, it will execute all the lines in LIFO order.
func main() {
	defer fmt.Println("World!!")
	defer fmt.Println("Hello ")
	fmt.Println("One")
	myDefer()
}

//The stack is of invidual functions, hence all the defer statements of this functions will be called after calling this function itself.
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
