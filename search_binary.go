package algorithms

// SearchBinary searches for a `needle` in `haystack`.
func SearchBinary(haystack []int, needle int) bool {
	var low int = -1
	var high int = len(haystack)
	var midpoint int

	for high-low > 1 {
		midpoint = high + low>>1

		if haystack[midpoint] > needle {
			high = midpoint
		} else if haystack[midpoint] < needle {
			low = midpoint
		} else {
			break
		}
	}

	if haystack[midpoint] == needle {
		return true
	} else {
		return false
	}
}
