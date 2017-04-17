// Package algorithms implements simple common programming algorithms
package algorithms

import "math"

// Finds the maximum sum of the elements of a subarray in a given array.
func MaxSubarray(array []int) int {
	maximum, current := .0, .0
	for i := 0; i < len(array); i++ {
		current = math.Max(.0, current+float64(array[i]))
		maximum = math.Max(maximum, current)
	}
	return int(maximum)
}
