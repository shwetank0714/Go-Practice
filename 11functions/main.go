package main

import "fmt"
func main()  {
	
	fmt.Println("Welcome to functions in Go Lang")

	sum,_,myStr := proAdderTwo(2,3,4,4,5,5,5)

	fmt.Println("The result of additon is:  ", sum, myStr)
}
// ...int sees the incomming values and makes it a slice
// func proAdder(values ...int) int {

// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}

// 	return total
// }
// can also return multiple values


func proAdderTwo(values ...int) (int, string,string) {

	total := 0
	for _, val := range values {
		total += val
	}

	return total, "total Value 1", "totalValueTwo"
}