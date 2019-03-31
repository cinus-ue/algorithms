package bucket_sort

func sortInBucket(bucket []int) {
	length := len(bucket)
	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		for j >= 0 && backup < bucket[j] {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = backup
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

func Sort(list []int) []int {

	num := len(list)
	max := max(list)
	buckets := make([][]int, num)

	index := 0
	for i := 0; i < num; i++ {
		index = list[i] * (num - 1) / max

		buckets[index] = append(buckets[index], list[i])
	}

	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(list[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

	return list
}
