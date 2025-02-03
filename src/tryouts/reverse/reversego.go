package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}

	n := len(arr)

	i := 0
	for i < n/2 {
		swap(&arr[i], &arr[n-i-1])
		i++
	}

	fmt.Println(arr)

	sort.Sort(arr)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
