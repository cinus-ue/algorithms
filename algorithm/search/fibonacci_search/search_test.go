package fibonacci_search

import (
	"fmt"
	"testing"
)

func TestFibonacciSearch(t *testing.T) {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sortedArray)
	x := 5
	fmt.Printf("fibonacci_search '%d': %d\n", x, FibonacciSearch(sortedArray, x))

}
