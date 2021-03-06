package util

import "testing"

func TestUtil(t *testing.T) {

	list := GetArrayOfSize(10)

	if len(list) != 10 {
		t.Error()
	}

	a := list[0]
	b := list[1]

	Swap(list, 0, 1)

	if a == list[0] || b == list[1] {
		t.Error()
	}
}
