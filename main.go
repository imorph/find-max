package main

import (
	"fmt"
	"sort"

	f "github.com/imorph/find-max/lib/fmax"
)

func main() {
	man := struct {
		Name string
		Age  int
	}

	people := []man {
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	man = f.Max(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("Man with longest name:", man)

	man = f.Max(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("Oldest man:", man)
}