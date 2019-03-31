package bubble_sort

func Sort(list []int) {

	for itemCount := len(list) - 1; ; itemCount-- {
		swap := false
		for i := 1; i <= itemCount; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
