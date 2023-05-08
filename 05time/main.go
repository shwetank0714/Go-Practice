package main

import (
	"fmt"
	"time"
)
// Example formats which time.Format excepts below: ---

// const dateFormat = "Monday"
// const dateFormatwithDay = "11-01-2006 Monday"
const currentTimeFormat = "15:04:05"

func main() {
	fmt.Println("Time study golang")

	presentTime := time.Now()

	fmt.Println(presentTime.Format(currentTimeFormat))

	// Creating some date manually -----------

	createdDate := time.Date(2023,time.March,14,0,45,45,0,time.UTC)
	fmt.Println("The created date is - ", createdDate)
}
