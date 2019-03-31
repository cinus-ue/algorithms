package selection_sort

import "github.com/cinus-ue/algorithms-go/algorithm/sort/utils"

func Sort(list []int) {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			utils.Swap(list, i, minIndex)
		}

	}
}
