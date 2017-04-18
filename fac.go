package algorithms

// Get factorial value of n.
func Fac(n int) int {
	if n < 1 {
		return 1
	} else {
		return n * Fac(n-1)
	}
}
