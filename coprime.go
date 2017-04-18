package algorithms

// Test whether a and b are coprime integers.
func Coprime(a, b int) bool {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a == 1
}
