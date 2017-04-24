package algorithms_test

import (
	"fmt"
	"testing"

	algorithms "github.com/yefremov/algorithms-with-go"
)

func ExampleSearchLinear() {
	fmt.Println(algorithms.SearchLinear([]int{1, 2, 3}, 2))
	fmt.Println(algorithms.SearchLinear([]int{1, 2, 3}, 4))
	// Output:
	// true
	// false
}

func TestSearchLinear(t *testing.T) {
	if algorithms.SearchLinear([]int{1, 2, 3}, 2) != true {
		t.Error("expected true")
	}

	if algorithms.SearchLinear([]int{1, 2, 3}, 4) != false {
		t.Error("expected false")
	}
}
