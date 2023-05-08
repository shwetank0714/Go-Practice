package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcome := "Welcome to User Input: "
	fmt.Println(welcome)
	fmt.Println("Please Enter some rating b/w 1 - 5: ")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	fmt.Println("Thanx for the rating: ", userInput)

	// Conversion of string to some meaning full Number Here:
	// theRatingValue, err := strconv.ParseFloat(userInput,64) this will create the error in the future 
	// if try to use the value returned by the functions because we are trying to lets say parse 4\n or 5\n to the 
	// Parse float function because we are reading string until \n is pressed 

	// so to avoid the addition of \n to the input we use the strings library which provide us various options 
	// to trim down the space and do many more operations on it!


	theRatingValue, err := strconv.ParseFloat(strings.TrimSpace(userInput),64)

	// If - Else block

	if err != nil {
		fmt.Println("There is some error -- ", err)
	} else {
		fmt.Println("Added one to the rating: --", theRatingValue + 1)
	}


}
