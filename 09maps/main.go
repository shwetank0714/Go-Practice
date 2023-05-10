package main

import (
	"fmt"
)

func main()  {
	
	fmt.Println("Maps ---------")

	laguagesMap := make(map[string]int)

	laguagesMap["JS"] = 2
	laguagesMap["C++"] = 21
	laguagesMap["Python"] = 22
	laguagesMap["Kotlin"] = 24
	laguagesMap["Go"] = 27

	fmt.Println(laguagesMap["JS"])

	// deleting in map
	delete(laguagesMap,"Kotlin")
	fmt.Println(laguagesMap)

	
	
}