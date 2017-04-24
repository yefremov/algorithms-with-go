package algorithms_test

import (
	"fmt"
	"testing"

	algorithms "github.com/yefremov/algorithms-with-go"
)

func ExampleFac() {
	for i := 0; i < 6; i++ {
		fmt.Println(algorithms.Fac(i))
	}
	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
}

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
