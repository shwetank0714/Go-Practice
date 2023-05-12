package main

import "fmt"

func main() {

	fruitList := []string{"Mango","Apple","Watermelon","Pineapple","Plums","Melon","Guava"}

	// Here the i gets the index of the elements in the slie over which we are iterating
	// kind of similar to the iterator in C++
	// for i := range fruitList {
	// 	fmt.Println(fruitList[i])
	// }

	// Similar to While loop
	starterVal := 0
	for starterVal < 5{
		fmt.Println(starterVal)
		starterVal++
	}
	


	//fmt.Print("\n\n")
	// for each kind of loop in go
	for _, value := range fruitList{
		if value == "Watermelon"{
			break
		} else{
			fmt.Println(value)
		}
	}
}