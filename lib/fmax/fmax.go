package fmax

import (
	"reflect"
)

// Max returns max value of provided slice. Panics if provided interface is not slice.
func Max(slice interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != Slice {
		panic(&ValueError{Method: "Swapper", Kind: v.Kind()})
	}
	return
}
