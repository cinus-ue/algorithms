package cocktail_sort

func Sort(list []int) {

	length := len(list)
	for i := 0; i < length/2; i++ {
		for j := i; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]

			}
		}

		for j := length - i - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}
