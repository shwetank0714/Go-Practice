package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices --")

	// Slice syntax
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "mango", "Pineapple")
	fmt.Println(fruitList)

	var integerList = []int{1,2,3,4,5,5,6,6}
	fmt.Println(integerList)

	// integerList[1:] == starting from 1 till end
	// integerList[1:n] == starting from 1 till n where n is not included
	integerList = append(integerList[1:])
	fmt.Println(integerList)
	integerList = append(integerList[1:4])
	fmt.Println(integerList)

	// Declaring slice dynamically

	highScores := make([]int, 5)
	highScores[0] = 1
	highScores[1] = 11
	highScores[2] = 111
	highScores[3] = 1111
	highScores[4] = 11111

	fmt.Println(highScores)

	// we can use the append in the slices even after allocating the memory to the slices
	// it will again re allocate the memory to the slice for the incoming values 
	highScores = append(highScores, 32,434,323)
	sort.Ints(highScores)
	fmt.Println(highScores)


	// Remove/Delete the elements from the slice

	var courses = []string{"python","C++","reactjs","Java","Golang","swift","ruby"}

	// deletion using append -------- Important
	// deleting the element from 3rd index
	var index = 3
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}