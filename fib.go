package algorithms

// Fib gets fibonacci value of `n`.
func Fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
