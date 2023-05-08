package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to User Input: "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Go:")

	// Comma, Ok Syntax || Error, ok syntax -------

	// ReadString('\n) will read the string until a \n  OR Enter is pressed  from the keyboard
	// (Succes, Errors) -> Error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("The Go rating is: ", input)
	fmt.Printf("The Type of this rating is: %T", input)
}
