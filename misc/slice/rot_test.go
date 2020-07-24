package slice

import (
	"reflect"
	"testing"
)

func TestRotLeft(t *testing.T) {

	intSlice := []int32{1, 2, 3, 4, 5}
	newSlice := RotLeft(intSlice, 2)
	want := []int32{3, 4, 5, 1, 2}

	if !reflect.DeepEqual(newSlice, want) {
		t.Errorf("RotLeft() = %v, want %v", newSlice, want)
	}

}
