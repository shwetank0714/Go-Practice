package main

import "fmt"

// struct definition

type User struct {
	Name   string
	Email  string
	Number int64
	Age    int64
	Status bool
}

func main() {

	myUser := User{Name: "Jacob",Email: "jacob@jacob.j", Number: 2324342441, Age: 25, Status: true}

	fmt.Println(myUser)
	// The below statement prints the struct as a whole Structure
	fmt.Printf("the Structure of myUser is %+v \n",myUser)
	fmt.Println(myUser.Age)
}
