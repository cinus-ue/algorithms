package ternary_search

import (
	"fmt"
	"testing"
)

func TestTernarySearch(t *testing.T) {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sortedArray)
	x := 11
	fmt.Printf("ternary_search '%d': %d\n", x, TernarySearch(sortedArray, x))
}
