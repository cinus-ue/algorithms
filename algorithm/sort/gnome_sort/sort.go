package gnome_sort

import "github.com/cinus-ue/algorithms-go/algorithm/sort/utils"

func Sort(list []int) {
	index := 0
	for index < len(list)-1 {
		if list[index] > list[index+1] {
			utils.Swap(list, index, index+1)
			if index != 0 {
				index -= 1
			}
		} else {
			index += 1
		}
	}

}
