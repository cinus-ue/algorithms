package insertion_search

import (
	"fmt"
	"testing"
)

func TestInsertionSearch(t *testing.T) {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sortedArray)
	x := 11
	fmt.Printf("insertion_search '%d': %d\n", x, InsertionSearch(sortedArray, x, 3, 7))
}
