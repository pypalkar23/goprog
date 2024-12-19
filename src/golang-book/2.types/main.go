package main

import "fmt"

func main() {
	fmt.Println("Hello\n\\'World")
	fmt.Println("1+1", 1+1)         //2
	fmt.Println("1.0+1.0", 1.0+1.0) //2 .0
	fmt.Println("1.0+1.0", 1.0+1)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])         // prints 101 i.e byte representation of the char
	fmt.Println(string("Hello World"[0])) // Prints H
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
