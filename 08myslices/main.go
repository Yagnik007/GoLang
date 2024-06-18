package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("These are slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Blackberry")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList) //Output: [Apple Tomato Peach] end limit non-inclusive

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 989
	highScores[2] = 433
	highScores[3] = 545
	// highScores[4] = 232 //Out of bound error

	highScores = append(highScores, 23423, 234, 324, 123, 3423)
	//It reallocate the memory for this slice and store this, IT'S WEIRD INIT?
	fmt.Println(highScores)

	sort.Ints(highScores) //Sorting

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	//We could use append method for removing the element as well
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	//Output: [reactjs javascript python ruby]
}
