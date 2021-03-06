package radix_sort

import (
	"fmt"
	"github.com/cinus-ue/algorithms/util"
	"testing"
)

func TestRadixSort(t *testing.T) {

	list := util.GetArrayOfSize(100)
	fmt.Println(list)
	RadixSort(list)
	fmt.Println(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}
