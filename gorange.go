package gorange

import "fmt"

func IntRange(first, last, step int) ([]int, error) {
	if last < first {
		msg := "last value must be greater than or equal to a first value"
		return nil, fmt.Errorf(msg)
	}
	sliceSize := 0
	if step > last {
		sliceSize = 1
	} else {
		sliceSize = (last - first + 1) / step
	}
	intRange := make([]int, sliceSize)
	value := first
	stepFactor := 0
	for i := 0; value <= last; i++ {
		stepFactor++
		intRange[i] = value
		value = first + step*stepFactor
	}
	return intRange, nil
}
