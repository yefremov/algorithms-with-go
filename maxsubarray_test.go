package algorithms_test

import (
	algorithms "github.com/yefremov/algorithms-with-go"
	"testing"
)

func TestMaxSubarray(t *testing.T) {
	if algorithms.MaxSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("expected 6")
	}
}
