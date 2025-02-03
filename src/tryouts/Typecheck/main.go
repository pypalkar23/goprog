package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

/*(func (p Person) Len() int {
	return Len(p)
}*/

func main() {
	people := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}

	arr := [...]int{1, 2, 3}

	fmt.Printf("Type of people: %s\n", reflect.TypeOf(people))
	fmt.Printf("Type of people: %s\n", reflect.TypeOf(arr))
}
