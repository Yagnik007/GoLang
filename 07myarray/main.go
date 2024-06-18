package main

import "fmt"

func main() {
	fmt.Println("This is an array!")

	var fruitList [4]string //We must specify the size of the array

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Strawberry"

	fmt.Println(fruitList) //Output:[Apple Banana  Strawberry]
	//There is a space in place of index 2 element, since we have not inserted any value at that particular index.

	fmt.Println(len(fruitList)) //Output: 4
}
