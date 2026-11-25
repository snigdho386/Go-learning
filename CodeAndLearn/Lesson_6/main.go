package main

import "fmt"

// Var is package level variable
// var myInt int = 10

// Group similar variables
// var (
// 	myInt    int    = 10
// 	myString string = "Hello World"
// )

func main() {

	// Explicit type assignment
	var myInt int = 10
	var myString string = "Hello World!"
	fmt.Println(myInt)
	fmt.Println(myString)

	// Implicit type assignment
	var age = 25
	fmt.Println(age)

	//Multiple variable decalration
	var year, name = 2023, "code & learn"
	fmt.Println(year, name)

	// Examples:
	// var apartmentNumber int         // Zero value to the variable
	// var apartmentNumber1 int = 2000 // When we want to assign a value to variable in declaration
	// var apartmentNumber2 = 2000
	// var apartmentNumber3, streetName = 2000, "Bayshore" // Multiple variable declaration on the same line
	// apartmentNumber4 := 2000
	// apartmentNumber5, streetstreetName := 2000, "Bayshore"

}
