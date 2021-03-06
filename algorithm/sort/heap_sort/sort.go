package heap_sort

import (
	"github.com/cinus-ue/algorithms/util"
)

func heapify(list []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && list[left] > list[largest] {
		largest = left
	}

	if right < n && list[right] > list[largest] {
		largest = right
	}

	if largest != i {
		util.Swap(list, i, largest)
		heapify(list, n, largest)
	}
}

func Sort(list []int) []int {

	for i := len(list)/2 - 1; i >= 0; i-- {
		heapify(list, len(list), i)
	}
	for i := len(list) - 1; i > 0; i-- {
		util.Swap(list, i, 0)
		heapify(list, i, 0)
	}
	return list
}
