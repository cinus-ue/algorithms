package heap_sort

import "github.com/cinus-ue/algorithms-go/algorithm/sort/utils"

func heapfy(list []int, n int, i int) {

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
		utils.Swap(list, i, largest)

		heapfy(list, n, largest)
	}
}


func Sort(list []int) []int {

	for i := len(list)/2 - 1; i >= 0; i-- {
		heapfy(list, len(list), i)
	}

	for i := len(list) - 1; i >= 0; i-- {
		utils.Swap(list, i, 0)
		heapfy(list, i, 0)
	}
	return list
}