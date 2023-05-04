package main

import "fmt"

// starting a const with CAPS letter means that the variable/value cab be access as a **PUBLIC**
const MyName = "Go Language"

func main() {

	// In Go lang unused varibles give error, and hence is not allowed to have

	var name string = "Go Lang"
	var number int = 38473
	var floatNumber float64 = 3293.232
	var booleanVar bool = true

	fmt.Println(name,number,floatNumber,booleanVar)

	// walrus operator ( := )for variables in golang
	// walrus operator must be used inside some method only, can't use globally
	newNumber :=  14
	newString := "new Name"
	newboolean := true
	fmt.Println(newNumber,newString,newboolean)

	// As default the go lang do not assign any garbage foa a data type
	// for Boolean -> default = false
	// for int, float -> default = 0
	// for string -> default = ""

	var newBool float64
	fmt.Println(newBool)
}