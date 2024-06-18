package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1
	fmt.Println("Random number between 1 and 6:", diceNumber)

	switch diceNumber{
		case 1: 
			fmt.Println("One")
		case 2: 
			fmt.Println("Two")
		case 3: 
			fmt.Println("Three")
			fallthrough //It will consider the next case as well
		case 4: 
			fmt.Println("Four")
		case 5: 
			fmt.Println("Five")
		case 6: 
			fmt.Println("Six")
		default: fmt.Println("Default") 
	}
}