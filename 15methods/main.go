package main

import "fmt"

func main() {
	fmt.Println("These are structs")
	//No inheritance in golang; No super or parent

	Yagnik := User{"Yagnik", "yagniktalaviya@example.co", true, 22}

	//This is how we access the values of a Struct
	fmt.Println(Yagnik)
	fmt.Printf("%+v\n", Yagnik)
	fmt.Printf("Email is %v and Age is %v\n", Yagnik.Email, Yagnik.Age)
	Yagnik.GetStatus()
	Yagnik.NewMail()
	fmt.Printf("Email is %v and Age is %v\n", Yagnik.Email, Yagnik.Age)
}

//All the variables need to be exported out from the file hence they are in uppercase, even the Struct Name is in uppercase obv. duh.

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// simpleVal int | This variable won't be exported - lowercase
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

//Passing the reference, this will change the value in the original User
func (u *User) NewMail() {
	u.Email = "dk@email.com"
	fmt.Println("New email is: ", u.Email)
}
