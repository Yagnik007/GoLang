package main

import "fmt"

func main() {
	fmt.Println("These are structs")
	//No inheritance in golang; No super or parent

	Yagnik := User{"Yagnik", "yagniktalaviya@example.co", true, 22}

	//This is how we access the values of a Struct
	fmt.Println(Yagnik)
	fmt.Printf("%+v\n", Yagnik)
	fmt.Printf("Email is %v and Age is %v", Yagnik.Email, Yagnik.Age)
}

//All the variables need to be exported out from the file hence they are in uppercase, even the Struct Name is in uppercase obv. duh.

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
