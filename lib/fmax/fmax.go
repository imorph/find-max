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

	for i := 0; i < v.Len(); i++ {
		fmt.Println("individual element of slece: ", v.Index(i))
	}

	fmt.Println("incominng slice is: ", v)

	return "mock"
}
