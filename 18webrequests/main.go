package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.notion.so/GoLang-a74ce0a8a4e34437a01a849744821fac"

func main() {
	fmt.Println("Web requests")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	//Output: Response is of type: *http.Response
	//This is a reference to the response, hence it would be easy to manipulate the response.

	defer response.Body.Close() //It is of utmost importance to close the connection and it is solely the responsibility of the caller!!!

	var buffer bytes.Buffer
	databytes, err := io.Copy(&buffer, response.Body)
	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
