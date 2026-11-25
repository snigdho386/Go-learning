package main

import (
	"fmt"
)

const AConst string = "This is a constant"

func main() {

	// Declaring vaiable in string
	// var aString string = "This is Go!"

	// fmt.Println(aString)

	// get the variable type
	// fmt.Printf("Type of aString: %T\n", aString)

	// var def float64
	// fmt.Printf("Type of def: %T\n", def)
	// fmt.Println(def)

	// NOTE : the := operator works only inside a function
	// If you declare a variable outside a function, you must use the var keyword

	// myString := aString
	// fmt.Println(myString)

	// IO stuffs
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text : ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("You entered:", input)
	// fmt.Printf("You entered: %T", input)

	// -- strconv package is used to convert string to other types
	// It's always a good idea to trim the input string before converting it to other types
	// The format for O/P of a func : variable,err if the function throws back an error.

	// fmt.Print("\nEnter a float : ")
	// input, _ = reader.ReadString('\n')
	// fmt.Println("You entered:", input)
	// afloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Println("Error parsing float:", err)
	// } else {
	// 	fmt.Println("You entered:", afloat)
	// 	fmt.Printf("You entered: %T", afloat)
	// }

	// Math stuffs
	// Type conversion is required to add float and int
	// a := 12.5
	// b := 1
	// c := a + float64(b)
	// fmt.Println(c)

	// i1, i2, i3 := 10, 20, 50

	// fmt.Print(math.Round(12.5555*1000) / 1000)

	// fmt.Println("\nSum of 3 numbers is : ", i1+i2+i3)
	// fmt.Printf("Average of 3 numbers is : %.3f", float64(i1+i2+i3)/3)

	// Date and time stuffs
	// n := time.Now()
	// fmt.Print("\nTime is: ", n)
	// t := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)
	// fmt.Println("\nTime t is: ", t)
	// fmt.Println((t.Format(time.ANSIC)))

	// parseTime, err := time.Parse(time.ANSIC, "Mon Jan 01 12:00:00 2024")

	// fmt.Println("Type of parsedTime is %T", parseTime)

	// if err != nil {
	// 	fmt.Println("Error parsing time:", err)
	// } else {
	// 	fmt.Println("Parsed time is:", parseTime)

	// }

	const name, dept = "GeeksforGeeks", "CS"

	// Calling Printf() function
	fmt.Printf("%s is a %s Portal.\n", name, dept)

}
