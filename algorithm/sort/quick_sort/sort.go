package quick_sort

import "github.com/cinus-ue/algorithms-go/algorithm/sort/utils"

func partitionRecursion(list []int, left int, right int) {

	low := left
	high := right
	pivot := list[(left+right)/2]
	for low <= high {
		for list[low] < pivot {
			low++
		}

		for list[high] > pivot {
			high--
		}
		if low <= high {
			utils.Swap(list, low, high)
			low++
			high--
		}
	}

	if left < high {
		partitionRecursion(list, left, high)
	}

	if low < right {
		partitionRecursion(list, low, right)
	}

}

func Sort(list []int) {
	partitionRecursion(list, 0, len(list)-1)
}
