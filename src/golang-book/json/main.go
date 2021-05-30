package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var titles []struct {
	Title string
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}

	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("Json marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var moviesList []Movie

	err1 := json.Unmarshal(data, &titles)
	if err1 != nil {
		fmt.Printf("Json marshalling failed: %s", err1)
	}
	fmt.Printf("%s\n", titles)

	err2 := json.Unmarshal(data, &moviesList)
	if err2 != nil {
		fmt.Printf("Json marshalling failed: %s", err2)
	}
	fmt.Println(moviesList)
}
