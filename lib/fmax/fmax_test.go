package fmax

import (
	"testing"
)

func TestMax001(t *testing.T) {
	inSlice := []string{"1", "2", "3", "42"}
	want := "42"
	got, err := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if err != nil {
		t.Errorf("Max() returned error: %q", err)
	}
	if got.(string) != want {
		t.Errorf("Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] ) == %q, want %q", got, want)
	}
}

func TestMax002(t *testing.T) {
	inSlice := []int{1, 2, 3, 42}
	want := 42
	got, err := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if err != nil {
		t.Errorf("Max() returned error: %q", err)
	}
	if got.(int) != want {
		t.Errorf("Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] ) == %q, want %q", got, want)
	}
}

func TestMax003(t *testing.T) {
	inSlice := []float32{1.001, 2.002, 3.003, 42.0042}
	var want float32 = 42.0042
	got, err := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if err != nil {
		t.Errorf("Max() returned error: %q", err)
	}
	if got.(float32) != want {
		t.Errorf("Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] ) == %f, want %f", got, want)
	}
}

func TestMax004(t *testing.T) {
	inSlice := []float32{}
	_, err := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if err.Error() != NilSlice {
		t.Errorf("Max() returned error: %q instead of %q", err, NilSlice)
	}
}

func TestMax005(t *testing.T) {
	inSlice := "not_a_slice"
	_, err := Max(inSlice, func(i, j int) bool { return inSlice[i] < inSlice[j] })
	if err.Error() != NotSlice {
		t.Errorf("Max() returned error: %q instead of %q", err, NotSlice)
	}
}