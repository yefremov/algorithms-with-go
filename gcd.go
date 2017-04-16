package algorithms

func Gcd(a, b int) int {
  if a % b == 0 {
    return b
  } else {
    return Gcd(b, a % b)
  }
}
