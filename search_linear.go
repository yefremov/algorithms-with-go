package algorithms

// SearchLinear searches for a `needle` in a `haystack`.
func SearchLinear(haystack []int, needle int) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
