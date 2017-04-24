package algorithms_test

import (
	"fmt"
	"testing"

	algorithms "github.com/yefremov/algorithms-with-go"
)

func ExampleMaxSub() {
	fmt.Println(algorithms.MaxSub([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// Output:
	// 6
}

func TestMaxSub(t *testing.T) {
	if algorithms.MaxSub([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("expected 6")
	}
}
