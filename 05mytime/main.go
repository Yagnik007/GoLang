package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is my time")

	//Present time values
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
	//This is exactly how the format is to get the current values of different attributes of time.

	//Any time values
	createdDate := time.Date(2023, time.December, 04, 12, 12, 12, 0, time.UTC)
	fmt.Println(createdDate)
	//This is a created date, year-month(this way particular)-day-hour-minute-second-timezone-utc
}
	