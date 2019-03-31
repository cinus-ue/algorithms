package fibonacci_search

const maxSize = 20

func FibonacciSearch(array []int, target int) int {
	length := len(array)

	low := 0
	high := length - 1

	var F [maxSize]int

	F[0] = 1
	F[1] = 1
	for i := 2; i < maxSize; i++ {
		F[i] = F[i-1] + F[i-2]
	}

	k := 0
	for length > F[k]-1 {
		k++
	}

	temp := make([]int, F[k]-1)
	for j := 0; j < length; j++ {
		temp[j] = array[j]
	}

	for i := length; i < F[k]-1; i++ {
		temp[i] = array[high]
	}
	for low <= high {
		mid := low + F[k-1] - 1
		if target < temp[mid] {
			high = mid - 1
			k = k - 1
		} else if target > temp[mid] {
			low = mid + 1
			k = k - 2
		} else if mid < length {
			return mid
		} else {
			return length - 1
		}
	}
	return -1
}
