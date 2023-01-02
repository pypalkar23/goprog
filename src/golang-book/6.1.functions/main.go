package main

import "fmt"

func main() {
	scores := [5]float64{10, 15, 563, 643, 345}
	fmt.Println(average(scores[:])) //passed it as a slice

	val1 := f2()
	fmt.Println(val1)

	val2, val3 := f3()
	fmt.Println(val2, val3)

	fmt.Println(add(24.0))
	fmt.Println(add(1.0, 4.0, 5.0))
	fmt.Println(add(scores[1:4]...)) //note ... when calling variadic function

	makeEven := makeEvenGenerator() //returns function
	fmt.Println(makeEven())         //0
	fmt.Println(makeEven())         //2
	fmt.Println(makeEven())         //4
	fmt.Println(factorial(uint(5)))

	deferDemonstration()
}

//	 name	param type	   out type
func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

// named returns
func f2() (r int) {
	r = 1
	return
}

//multiple returns
func f3() (int, int) {
	return 5, 6
}

//multiple named returns
//can collapse the params with same type for e.g (d1,d2 int)
func f4() (d1 int, d2 int) {
	d1 = 10
	d2 = 15
	return
}

//variadic functions
func add(args ...float64) float64 {
	total := 0.0
	for _, val := range args {
		total += val
	}
	return total
}

//closure
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}

}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func deferDemonstration() {
	defer second()
	first()
}
