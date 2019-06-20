package main

import (
	"fmt"
	"log"

	f "github.com/imorph/find-max/lib/fmax"
)

func main() {
	type man struct {
		Name string
		Age  int
	}
	people := []man{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	bigMan, err := f.Max(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	if err != nil {
		log.Println("Max() returned error: ", err)
	} else {
	fmt.Println(`Man with "longest" name:`, bigMan)
	}

	bigMan, err = f.Max(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	if err != nil {
		log.Println("Max() returned error: ", err)
	} else {
	fmt.Println("Oldest man:", bigMan)
	}
}
