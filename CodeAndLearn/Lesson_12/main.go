package main

import (
	"fmt"
	"strconv"
)

// In this lesson we will cover
// 1. OR, AND, XOR AND NOT and the INVERT operator
// 2. Printing ints in binary forms

func main() {

	var a uint8 = 202
	var b uint8 = 141

	// The OR operator
	c := a | b
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	// The XOR operator
	c = a ^ b
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	// The AND NOT or Bit clear operator
	c = a &^ b
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	// The INVERT or the NOT operator
	c = ^a
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

}
