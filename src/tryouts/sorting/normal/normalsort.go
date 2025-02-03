package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"e", "a", "b", "c", "d"}

	sort.Strings(arr)
	fmt.Println(arr)

	sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	fmt.Println(arr)
}
