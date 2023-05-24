package main

import "fmt"

func main() {

	// By adding the defer keyword the functions are exectuted at the end of the function inside which they are called in
	// And they are called in the reverse order. (LIFO) approach
	defer fmt.Println("Line 1")
	fmt.Println("Hello Defers")
	defer fmt.Println("Line 2")
	defer fmt.Println("Line 3")
	loopDefer()
}

func loopDefer(){

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// OutPut Predict:
// Hello Defers
// 4
// 3
// 2
// 1
// 0
// Line3
// Line2
// Line1