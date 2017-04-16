package algorithms_test

import (
  "testing"
  algorithms "github.com/yefremov/algorithms-with-go"
)

func TestFib(t *testing.T) {
  if algorithms.Fib(1) != 1 {
    t.Error("expected 1")
  }

  if algorithms.Fib(3) != 3 {
    t.Error("expected 3")
  }

  if algorithms.Fib(7) != 21 {
    t.Error("expected 21")
  }
}

func BenchmarkFib(b *testing.B) {
    for i := 1; i < b.N; i++ {
        algorithms.Fib(i)
    }
}
