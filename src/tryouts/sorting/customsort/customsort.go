package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

	track := make(map[string]string)

	track["mandar"] = "mandar"
	track["dipesh"] = "dipesh"

	for _, person := range people {
		fmt.Println(person.Name)
	}

	for key, value := range track {
		fmt.Printf("%s, %s\n", key, value)
	}
}
