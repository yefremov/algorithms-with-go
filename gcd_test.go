package algorithms_test

import (
  "testing"
  algorithms "github.com/yefremov/algorithms-with-go"
)

func TestGcd(t *testing.T) {
  if algorithms.Gcd(468, 24) != 12 {
    t.Error("expected 12")
  }

  if algorithms.Gcd(24, 12) != 12 {
    t.Error("expected 12")
  }

  if algorithms.Gcd(135, 19) != 1 {
    t.Error("expected 1")
  }

  if algorithms.Gcd(19, 2) != 1 {
    t.Error("expected 1")
  }

  if algorithms.Gcd(2, 1) != 1 {
    t.Error("expected 1")
  }
}

func BenchmarkGcd(b *testing.B) {
    for i := 1; i < b.N; i++ {
        algorithms.Gcd(b.N - i, i)
    }
}
