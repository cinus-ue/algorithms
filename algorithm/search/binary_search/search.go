package binary_search

func NonRecursionBinarySearch(array []int, target int) int {
	low, high := initSearch(array)
	for low <= high {
		middle := (low + high) / 2
		if array[middle] == target {
			return middle
		} else if array[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

func RecursionBinarySearch(array []int, target int) int {
	low, high := initSearch(array)
	return recursionSearch(array, target, low, high)
}

func recursionSearch(array []int, target int, low int, high int) int {
	middle := (low + high) / 2
	if array[middle] == target {
		return middle
	} else if array[middle] < target {
		return recursionSearch(array, target, middle+1, high)
	} else {
		return recursionSearch(array, target, low, middle-1)
	}

}

func initSearch(array []int) (int, int) {
	if array == nil {
		panic("Nil array.")
	}
	low := 0
	high := len(array) - 1
	if high < low {
		panic("Empty array.")
	}
	return low, high
}
