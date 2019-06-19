package fmax

import (
	"reflect"
	"log"
	"fmt"
)

// Max returns max value of provided slice. Panics if provided interface is not slice.
func Max(slice interface{}, comparator func(i, j int) bool) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		log.Println("object is not a slice, it is: ", v.Kind())
		panic("AAAAA! I am in panic!")
	}

    fmt.Println("incominng slice is: ", v)

	return "mock"
}
