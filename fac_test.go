package algorithms_test

import (
  "testing"
  algorithms "github.com/yefremov/algorithms-with-go"
)

func TestFac(t *testing.T) {
  if algorithms.Fac(0) != 1 {
    t.Error("expected 1")
  }

  if algorithms.Fac(3) != 6 {
    t.Error("expected 6")
  }

  if algorithms.Fac(5) != 120 {
    t.Error("expected 120")
  }
}

func BenchmarkFac(b *testing.B) {
    for i := 1; i < b.N; i++ {
        algorithms.Fac(i)
    }
}
