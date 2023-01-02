package main

import "fmt"

func main() {
	//Go has only for(no while, do , until , foreach) loop which can
	//for variant 1 similar to while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	//variant 2 usual one j is not available outside loop
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	//if else
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	//switch
	var input int //32 bit integer equivalent to int32
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &input)
	cal := input % 2

	switch cal {
	case 1:
		fmt.Println("odd")
		fmt.Println("to check multiline switch")
	case 0:
		fmt.Println("even")
	}
}
