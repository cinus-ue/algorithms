package comb_sort

func Sort(list []int) {

	gapValue := len(list)
	swapCount := 1
	for gapValue >= 1 && swapCount != 0 {
		if gapValue != 1 {
			gapValue = int(float64(gapValue) / float64(1.3))
		}
		swapCount = 0
		firstItem := 0
		secondItem := gapValue
		for secondItem != len(list) {
			if list[firstItem] > list[secondItem] {
				list[firstItem], list[secondItem] = list[secondItem], list[firstItem]
				swapCount += 1
			}
			firstItem += 1
			secondItem += 1
		}
	}

}
