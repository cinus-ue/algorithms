package merge_sort

func Sort(list []int) {
	if len(list) < 2 {
		return
	}
	middle := len(list) / 2
	var s = make([]int, middle+1)
	Sort(list[:middle])
	Sort(list[middle:])
	if list[middle-1] <= list[middle] {
		return
	}

	copy(s, list[:middle])
	l, r := 0, middle
	for i := 0; ; i++ {
		if s[l] <= list[r] {
			list[i] = s[l]
			l++
			if l == middle {
				break
			}
		} else {
			list[i] = list[r]
			r++
			if r == len(list) {
				copy(list[i+1:], s[l:middle])
				break
			}
		}
	}
	return
}
