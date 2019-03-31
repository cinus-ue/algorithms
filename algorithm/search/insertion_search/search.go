package insertion_search

func InsertionSearch(array []int, value int, low int, high int) int {
	mid := low + (value-array[low])/(array[high]-array[low])*(high-low)
	if array[mid] == value {
		return mid
	}
	if array[mid] > value {
		return InsertionSearch(array, value, low, mid-1)
	}

	if array[mid] < value {
		return InsertionSearch(array, value, mid+1, high)
	}
	return -1
}
