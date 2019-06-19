package fmax

import (
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
	var biggest int
	for i := 1; i < v.Len(); i++ {
		if comparator(i-1, i) {
			biggest = i
		}
	}
	return v.Index(biggest)
}
