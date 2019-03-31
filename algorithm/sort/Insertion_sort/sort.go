package Insertion_sort

func Sort(list []int) {
	for i := 1; i < len(list); i++ {
		tmp := list[i]
		j := i - 1
		for j > 0 && tmp < list[j-1] {
			list[j] = list[j-1]
			j--
		}
		list[j] = tmp
	}
}
