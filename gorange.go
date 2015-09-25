package gorange

import "fmt"

func IntRange(first, last, step int) ([]int, error) {
	var intRange []int
	if last < first {
		msg := "last value must be greater than or equal to a first value"
		return intRange, fmt.Errorf(msg)
	}
	value := first
	stepFactor := 0
	for value <= last {
		stepFactor++
		intRange = append(intRange, value)
		value = first + step*stepFactor
	}
	return intRange, nil
}
