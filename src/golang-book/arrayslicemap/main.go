package main

import "fmt"

func main() {
	//array declaration
	var x [5]int
	fmt.Println(x) //prints [0 0 0 0 0]
	x[4] = 100     //notice no := here as array was initialized in the first step
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)
	//fmt.Println(total / len(y)) -> mismatched types as total is float and len(y) is int
	fmt.Println(total / float64(len(y)))

	//another for loop variation for array iteration
	for i, value := range y {
		fmt.Println(i, value)
	}

	//if dont want to use i
	for _, value := range y {
		fmt.Println(value)
	}
}
