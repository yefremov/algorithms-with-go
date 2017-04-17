// Package algorithms implements simple common programming algorithms
package algorithms

// Linear search of haystack for a needle.
func SearchLinear(haystack []int, needle int) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
