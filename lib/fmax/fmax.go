package fmax

import (
	"fmt"
	"log"
	"reflect"
)

// Max returns max value of provided slice. Panics if provided interface is not slice.
func Max(slice interface{}, comparator func(i, j int) bool) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		log.Println("object is not a slice, it is: ", v.Kind())
		panic("AAAAA! I am in panic!")
	}

	var n, biggest int
	for i := 1; i < v.Len(); i++ {
		fmt.Println("individual element of slice: ", v.Index(i))
		if comparator(i-1, i) {
			n = i
			biggest = n
		}
	}
	fmt.Println("incominng slice is: ", v)
	fmt.Println("biggest element is: ", v.Index(biggest))
	return v.Index(biggest)
}
