package algorithms

// Coprime test whether `a` and `b` are relatively prime integers.
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
