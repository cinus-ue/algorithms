package counting_sort

func Sort(list []int) {
	maxValue := max(list)
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(list)

	for i := 0; i < length; i++ {
		bucket[list[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			list[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

}

func max(list []int) int {
	max := list[0]
	for _, e := range list {
		if e > max {
			max = e
		}
	}
	return max
}
