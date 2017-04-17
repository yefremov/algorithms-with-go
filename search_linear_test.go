package algorithms_test

import (
	algorithms "github.com/yefremov/algorithms-with-go"
	"testing"
)

func TestSearchLinear(t *testing.T) {
	if algorithms.SearchLinear([]int{1, 2, 3}, 2) != true {
		t.Error("expected true")
	}

	if algorithms.SearchLinear([]int{1, 2, 3}, 4) != false {
		t.Error("expected false")
	}
}
