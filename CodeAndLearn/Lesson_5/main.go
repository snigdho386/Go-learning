package main

import "fmt"

func main() {
	var myString string

	// zero value for string is ""
	// fmt.Println(myString)
	myString = "Hello World"
	fmt.Println(myString)

	myString = `Welcome to 
Golang course`

	fmt.Println(myString)
	var firstName, lastName string
	firstName = "code"
	lastName = "learn"

	var fullName string = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
