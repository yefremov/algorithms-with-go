package algorithms_test

import (
	"fmt"
	"reflect"
	"testing"

	algorithms "github.com/yefremov/algorithms-with-go"
)

func ExampleSortBubble() {
	fmt.Println(algorithms.SortBubble([]int{3, 3, 7, 1}))
	// Output:
	// [1 3 3 7]
}

func TestSortBubble(t *testing.T) {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if reflect.DeepEqual(algorithms.SortBubble(a), b) != true {
		t.Error("expected true")
	}

	c := []int{6, 2, 6, 3, 2, 3}
	d := []int{2, 2, 3, 3, 6, 6}

	if reflect.DeepEqual(algorithms.SortBubble(c), d) != true {
		t.Error("expected true")
	}
}
