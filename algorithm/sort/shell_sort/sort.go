package shell_sort

func Sort(list []int) {
	length := len(list)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := list[i]
			j := i - gap
			for j >= 0 && list[j] > temp {
				list[j+gap] = list[j]
				j -= gap
			}
			list[j+gap] = temp
		}
		gap = gap / 3
	}
	return
}
