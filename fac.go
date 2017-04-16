package algorithms

func Fac(n int) int {
  if n < 1 {
    return 1
  } else {
    return n * Fac(n - 1)
  }
}
