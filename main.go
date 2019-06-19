package main

import (
	"fmt"

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
	bigMan := f.Max(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println(`Man with "longest" name:`, bigMan)
	bigMan = f.Max(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("Oldest man:", bigMan)

	ints := []int{1, 2, 4, 10, 14, 3, 42}
	floats := []float32{1.23, 2.01, 4.3, 10.2, 14.7, 3.05, 42.00001}
	fmt.Println("This is slice of ints: ", ints)
	fmt.Println("This is biggest element: ", f.Max(ints, func(i, j int) bool { return ints[i] < ints[j] }))
	fmt.Println("This is slice of floats: ", floats)
	fmt.Println("This is biggest element: ", f.Max(floats, func(i, j int) bool { return floats[i] < floats[j] }))
}
