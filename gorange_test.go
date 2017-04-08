package gorange

import (
	"reflect"
	"testing"
)

func TestIntRange(t *testing.T) {
	validSlice := []int{1, 3, 5, 7, 9}

	resultSlice, err := IntRange(1, 10, 2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(resultSlice, validSlice) {
		t.Errorf("Expected %v got %v\n", validSlice, resultSlice)
	}

	validSlice = []int{1}
	resultSlice, err = IntRange(1, 100, 100)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(resultSlice, validSlice) {
		t.Errorf("Expected %v got %v\n", validSlice, resultSlice)
	}

	validSlice = []int{5, 6, 7, 8, 9}
	resultSlice, err = IntRange(5, 9, 1)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(resultSlice, validSlice) {
		t.Errorf("Expected %v got %v\n", validSlice, resultSlice)
	}

	validSlice = []int{1, 3, 5, 7, 9, 11}
	resultSlice, err = IntRange(1, 11, 2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(resultSlice, validSlice) {
		t.Errorf("Expected %v got %v\n", validSlice, resultSlice)
	}

	validSlice = []int{1, 4, 7, 10}
	resultSlice, err = IntRange(1, 10, 3)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(resultSlice, validSlice) {
		t.Errorf("Expected %v got %v\n", validSlice, resultSlice)
	}
}
