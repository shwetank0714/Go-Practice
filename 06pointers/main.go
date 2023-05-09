package main

import "fmt"

func getSum(firstNum int, secNum int, resultVarPointer *int) {
	fmt.Println("Value in the pointer variable - ", resultVarPointer)
	*resultVarPointer = firstNum + secNum
}

func main() {
	fmt.Println("Welcome to pointers  (*)")

	var resultVar int
	getSum(45,89, &resultVar)

	fmt.Println("The sum of two numbers is: ", resultVar)

	tempNumber := 45
	var ptr = &tempNumber

	fmt.Printf("The type of variable ptr is- %T \n", ptr)
	fmt.Println("The value inside pointer- ",*ptr)
	fmt.Println("The value of pointer- ", ptr)

	// Arthematics on pointers
	*ptr = *ptr + 75
	fmt.Println("The value inside pointer- ",*ptr)
	*ptr = *ptr * 5
	fmt.Println("The value inside pointer- ",*ptr)
	*ptr = *ptr - 11
	fmt.Println("The value inside pointer- ",*ptr)
	*ptr = *ptr / 3
	fmt.Println("The value inside pointer- ",*ptr)

}