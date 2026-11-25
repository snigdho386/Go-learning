package main

import "fmt"

const (
	x                     = 10
	y               int32 = 15
	applicationName       = "Lesson 7"
)

func main() {

	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = x
	fmt.Println(b)

	var c int
	c = int(y)
	fmt.Println(c)

	var d int32
	d = y
	fmt.Println(d)

	// only complex, real, imag, len and cap
	const z = complex(10.2, 100.9)
	const l = imag(z)

}
