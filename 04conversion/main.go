package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizzza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//strconv is used to convert strings to other formats, while converting (Parsing) we'll encounter an error where the conversion will have a "\n" at the end and hence we'll need to trim the space out of input.
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your ratin: ", numRating+1)
	}
}