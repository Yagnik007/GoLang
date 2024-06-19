package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//creation: os
//reading: ioutils
//manipulation: ioutils

func main() {
	fmt.Println("Files")
	content := "This will go into a file."

	//creating a file.
	file, err := os.Create("./myTxtFile")

	checkNilError(err)

	//Creating a file will return the length of the file
	//We will require io and os package for reading and writing the file.
	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println(length)
	//This is the way to close the file.
	defer file.Close()
	readFile("./myTxtFile")
}

//The format that is being read from the file is in the format of byte.

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println(databyte)
	//Output: [84 104 105 115 32 119 105 108 108 32 103 111 32 105 110 116 111 32 97 32 102 105 108 101 46]
	fmt.Println(string(databyte))
}

//Using function to check for nil error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
