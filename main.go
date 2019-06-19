package main

import (
	//"fmt"

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

	//var iAmInt int
	//bigMan := f.Max(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	//fmt.Println("Man with longest name:", bigMan)

	//bigMan = f.Max(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	//fmt.Println("Oldest man:", bigMan)

	_ = f.Max(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	// panic
	//_ = f.Max(iAmInt, func(i, j int) bool { return people[i].Age < people[j].Age })
}
