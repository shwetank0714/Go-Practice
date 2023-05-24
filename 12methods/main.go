package main

import "fmt"



type User struct {
	Name   string
	Email  string
	Number int64
	Age    int64
	Status bool
}

// method
func (u User) GetStatus() {
	if !u.Status{
		fmt.Println("Not Online")
	} else{
		fmt.Println("User Online")
	}
}

func (u User) changeMailShallowCopy(mailString string){
	u.Email = mailString
	fmt.Println("Email changed successfully: ", u.Email)
}
func (u *User) changeMailDeepCopy(mailString string){
	u.Email = mailString
	fmt.Println("Email changed successfully: ", u.Email)
}

func main() {

	myUser := User{Name: "Jacob", Email: "jacob@jacob.j", Number: 2324342441, Age: 25, Status: true}

	fmt.Println(myUser)
	// The below statement prints the struct as a whole Structure
	fmt.Printf("the Structure of myUser is %+v \n", myUser)
	fmt.Println(myUser.Age)

	myUser.GetStatus()
	fmt.Println("my email: ", myUser.Email)
	myUser.changeMailShallowCopy("helloworld@go.dev.com")
	fmt.Println("my email: ", myUser.Email)
	myUser.changeMailDeepCopy("newHelloWorld@go.dev")

	fmt.Println("my email: ", myUser.Email)

}
