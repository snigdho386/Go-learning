package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hey there!")

	// Variable of type uint8
	var smallPositiveValue uint8 = 255
	fmt.Println(smallPositiveValue)

	//variable of type int8
	var smallPositiveNegative int8
	smallPositiveNegative = -10

	fmt.Println(smallPositiveNegative)

	// Typecasting -> <TARGET_TYPE>(VARIABLE_NAME)
	var myInt int

	myInt = int(smallPositiveNegative)

	fmt.Printf("%t", myInt)

	var myByte byte = 'A'
	fmt.Println(myByte)
	// byte is mainly used to represent raw data
	// and also to distinguish between uint8 and byte
	// since byte is an alias for uint8

	// Rune is an alias for int32
	var myRune rune = 'B'
	fmt.Println(myRune)

}
