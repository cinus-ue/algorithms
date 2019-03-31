package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sortedArray)
	x := 11
	fmt.Printf("RecursionBinarySearch '%d': %d\n", x, RecursionBinarySearch(sortedArray, x))
	fmt.Printf("NonRecursionBinarySearch '%d': %d\n", x, NonRecursionBinarySearch(sortedArray, x))
}
