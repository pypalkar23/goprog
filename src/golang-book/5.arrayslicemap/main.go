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

	//shorthand syntax for array declaration
	z := [5]float64{12, 25, 265, 56, 165}
	fmt.Println(z)

	//for arrays that does not fit in one line
	t := [5]float64{
		93,
		12,
		12, //this extra comma is needed as per syntax
	}

	fmt.Println(t)

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

	///---Slices Section Start
	//empty slice-seems useless lets see
	var slice1 []float64
	fmt.Println(slice1)

	//slice with size 5
	slice2 := make([]float64, 5)
	fmt.Println(slice2)

	//slice with size 5 capacity 10
	slice3 := make([]float64, 5, 10)
	slice3[0] = 10
	fmt.Println(slice3, len(slice3))

	//slice created from array
	arr := [5]float64{1, 2, 3, 4, 5}
	slice4 := arr[0:4] //inclusive of start and end

	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
	}
	//array and slice have same capacity and slice
	fmt.Printf("Cap of the array: %d\n", cap(arr))
	fmt.Printf("Cap of the slice: %d\n", cap(slice4))
	fmt.Printf("Length of the array: %d\n", len(arr))

	//appending new elements to a slice which is already full.
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 15)

	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
	}

	/*slice gets bigger than the array underneath, new slice gets created (doubled in capacity) with new array underneath,
	and existing elements get copied to it.*/

	fmt.Printf("Cap of the slice: %d\n", cap(slice4))
	fmt.Printf("Cap of the array: %d\n", cap(slice4))
	fmt.Printf("Length of the array: %d\n", len(arr))

	//append slice
	slice5 := []int{1, 2, 3}
	slice6 := append(slice5, 4, 5)
	fmt.Println(slice5, slice6) // slice6-> 1,2,3,4,5

	slice7 := []int{1, 2, 3}
	slice8 := make([]int, 2)
	copy(slice7, slice8) // will copy only first two elements as slice8 has size 2
	fmt.Println(slice1, slice2)

	//--slice end
	//-------maps start---------
	var map1 map[string]int     //key->string value->int note:this is uninitialized
	map1 = make(map[string]int) //initialized it here
	map1["key"] = 10
	fmt.Println(map1)

	delete(map1, "key")
	fmt.Println(map1)

}
