package main

import (
	"fmt"
	"math"
)

/*
	type Circle Struct{
		x float64,
		t float64,
		r float64
	}
*/

func main() {
	var c1 Circle
	c1 = Circle{x: 0, y: 0, r: 5}
	fmt.Println(c1) //{0,0,5}

	c1.x = -2
	c1.y = -10

	//if order of the field is known field names can be omitted
	c2 := Circle{0, 0, 5}
	fmt.Println(c2)

	c3 := new(Circle) //allocate memory for Circle struct instance returns pointer to the c3
	fmt.Println(c3)   // &{0,0,0}
	fmt.Println(*c3)  //{0,0,0}

	c3.x = 10
	c3.y = 10
	c3.r = 10
	fmt.Println(*c3)

	fmt.Println(func_calcArea(c1))
	fmt.Println(c2.method_area())
	fmt.Println(c3.method_area())

	p := new(Person)
	p.name = "Mandar"
	p.talk()

	android := new(Android)
	android.P = *p //since new returns pointer
	android.P.talk()
	p.name = "Other Name"
	android.P.talk() // will still print Mandar that p was passed as value not reference

	android1 := Android1{}
	android1.Person = *p
	android1.Person.talk()
	android1.talk() //since android1's embedded type is not named we can call the talk directly on android
}

type Circle struct {
	x, y, r float64
}

//normal function that takes circle instance
func func_calcArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//method declaration
//    receiver
func (c *Circle) method_area() float64 {
	return math.Pi * c.r * c.r //notice no *c reference here
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hello, my name is", p.name)
}

type Android struct {
	P     Person //Embedded Type
	Model string
}

type Android1 struct {
	Person //another example of Embedded type with no instance name
	Model  string
}
