package fmax

import (
	"reflect"
	"errors"
)

const (
	// NotSlice error string returned when given object is not slice
	NotSlice = "not_a_slice_error"
	// NilSlice error string returned when given slice is nil
	NilSlice = "nil_slice_error"
)

// Max returns max value of provided slice. Returns error when it is not slice or nil slice 
func Max(slice interface{}, comparator func(i, j int) bool) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		err := errors.New(NotSlice)
		return nil, err
	}

	vLen := v.Len()
	if vLen == 0 {
		err := errors.New(NilSlice)
		return nil, err
	}

	var biggest int
	for i := 1; i < vLen; i++ {
		if comparator(i-1, i) {
			biggest = i
		}
	}

	return v.Index(biggest).Interface() , nil
}
