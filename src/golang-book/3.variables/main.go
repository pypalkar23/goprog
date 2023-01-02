package main

import (
	"fmt"
)

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

	const (
		tmp = 1
		tmp2
		tmp3 = 2
		tmp4
	)
	fmt.Println(tmp, tmp2, tmp3, tmp4) //1 1 2 2
	//value similar to the previous variable gets assigned to the current variable if value is not specified explicitly.

	type Weekday int
	const (
		Monday Weekday = iota //iota generator every next constant declared gets value which is one greater than previous starting from 0
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		_   = 1 << (10 * iota) // 1 << 0 = 1
		KiB                    // 1 << (10*1)= 1024
		MiB                    // 1 << (10*2)= 1048576
	)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
