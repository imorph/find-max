package fmax

import (
	"testing"
	"reflect"
)

func TestMax001(t *testing.T) {
	inSlice := []string{"1","2","3","42"}
	want := "42"
	got := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] ) == %q, want %q", got, want)
	}
}
