package main

import "fmt"

func main () {
	fmt.Println("Welcome to arrays in go lang")

	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "Pineapple"
	
	fruitList[3] = "Mango"
	fruitList[4] = "Orange"
	
	

	fmt.Println("The fruit list is: ", fruitList)

	vegList := [5]string{"Spinach","Okra","Brinjal","beans","JackFruit"}
	var numList = [5]int{1,2,3,4,5}

	fmt.Println("Two vegetable lists are: ")
	fmt.Println(vegList)
	fmt.Println(numList)

}