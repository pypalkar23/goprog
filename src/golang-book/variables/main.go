package main

import "fmt"

func main() {
	var x string = "Hello World" // declaration & assignment in one stmt
	fmt.Println(x)

	var y string // declaration
	y = "first"  // assignment
	fmt.Println(y)

	y = "second" //  reassigning
	fmt.Println(y)

	y = y + "first" //concatenate
	fmt.Println(y)
	fmt.Println(x == y) // check if equal -> false

	z := "third" //shorthand for variable declaration type is considered implictly by compilers
	w := 5       //another example for shorthand declaration
	w += 1
	fmt.Println(z, w)

	const pi float64 = 3.14 //constant variable
	fmt.Println(pi)

	var (
		a = 5
		b = 10
		c = 15
	) //multiple variable declaration

	fmt.Println(a, b, c)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
