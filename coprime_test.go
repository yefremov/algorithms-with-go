package algorithms_test

import (
	algorithms "github.com/yefremov/algorithms-with-go"
	"testing"
)

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
