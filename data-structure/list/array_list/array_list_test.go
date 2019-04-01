package array_list

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {

	l := New()

	if !l.IsEmpty() || l.Length != 0 {
		t.Error()
	}

	l.Add(0)
	l.Add(1)
	l.Add(2)

	if l.Length != 3 {
		t.Error()
	}

	index1, _ := l.IndexOf(1)
	if index1 != 1 {
		fmt.Println(index1)
		t.Error()
	}

	if l.Contains(3) {
		t.Error()
	}

	if l.Get(0) != 0 || l.Get(1) != 1 || l.Get(2) != 2 {
		fmt.Println(l.Get(0))
		fmt.Println(l.Get(1))
		fmt.Println(l.Get(2))
		t.Error()
	}

	l.Set(0, 3)
	if l.Get(0) != 3 {
		t.Error()
	}

	l.Remove(0)
	if l.Length != 2 {
		t.Error()
	}

	l.Clear()
	if !l.IsEmpty() {
		t.Error()
	}

}
