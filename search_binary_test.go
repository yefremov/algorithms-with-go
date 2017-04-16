package algorithms_test

import (
  "testing"
  algorithms "github.com/yefremov/algorithms-with-go"
)

func TestSearchBinary(t *testing.T) {
  if algorithms.SearchBinary([]int {1,2,3}, 2) != true {
    t.Error("expected true")
  }

  if algorithms.SearchBinary([]int {1,2,3}, 4) != false {
    t.Error("expected false")
  }
}
