package gorange

import "fmt"

func IntRange(first, last int) ([]int, error) {
	var intRange []int
	if last < first {
		msg := "last value must be greater than or equal to a first value"
		return intRange, fmt.Errorf(msg)
	}
	for i := first; i <= last; i++ {
		intRange = append(intRange, i)
	}
	return intRange, nil
}
