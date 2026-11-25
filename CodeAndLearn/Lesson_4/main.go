package main

import "fmt"

func main() {
	var smallFloat float32
	fmt.Println(smallFloat)

	smallFloat = 23.0234324
	fmt.Println(smallFloat)

	var bigFloat float64
	bigFloat = 43.842492734297

	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)

	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)

	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPart)

}
