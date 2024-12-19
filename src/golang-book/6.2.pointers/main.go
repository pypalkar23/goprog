package main

import "fmt"

type Car struct {
	model string
}

func (c *Car) PrintModel() {
	fmt.Println(c.model)
}

func main() {
	i := 72
	iPtr := &i //&returns the pointer
	changeByReference(iPtr)
	fmt.Println(i)

	//creates enough memory to fit value for that type and returns pointer to it
	//useful for creating pointer without initializing variable
	iPtr2 := new(int)
	changeByReference(iPtr2)

	y := 14.08
	squareByPointer(&y)
	fmt.Println(y)

	c := Car{model: "Toyota Corolla LE"}
	cPtr := &c
	cPtr.PrintModel()
}

func changeByReference(ptr *int) {
	//only ptr=1 means changing the address which isnt allowed in go
	//*ptr means access the value referenced by pointer
	*ptr = 1

}

func squareByPointer(x *float64) {
	*x = *x * *x
}
