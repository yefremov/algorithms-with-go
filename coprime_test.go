package algorithms_test

import (
	"fmt"
	"testing"

	algorithms "github.com/yefremov/algorithms-with-go"
)

func ExampleCoprime() {
	fmt.Println(algorithms.Coprime(1, 1))
	fmt.Println(algorithms.Coprime(2, 2))
	fmt.Println(algorithms.Coprime(13, 27))
	fmt.Println(algorithms.Coprime(20536, 7826))
	// Output:
	// true
	// false
	// true
	// false
}

func TestCoprime(t *testing.T) {
	if algorithms.Coprime(1, 1) != true {
		t.Error("expected true")
	}

	if algorithms.Coprime(1, 2) != true {
		t.Error("expected true")
	}

	if algorithms.Coprime(2, 2) != false {
		t.Error("expected false")
	}

	if algorithms.Coprime(13, 27) != true {
		t.Error("expected true")
	}

	if algorithms.Coprime(20536, 7826) != false {
		t.Error("expected false")
	}
}

func BenchmarkCoprime(b *testing.B) {
	for i := 1; i < b.N; i++ {
		algorithms.Coprime(i, i)
	}
}
