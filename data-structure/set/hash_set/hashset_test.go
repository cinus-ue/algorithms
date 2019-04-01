package hash_set

import (
	"fmt"
	"testing"
)

func TestHashSet(t *testing.T) {

	s := New()
	if !s.IsEmpty() {
		t.Error()
	}

	s.Add(1, 2, 3, 4)
	if s.Size() != 4 {
		fmt.Println(s.Size())
		t.Error()
	}

	s.Remove(2)
	if s.Size() != 3 {
		fmt.Println(s.Size())
		t.Error()
	}

	if !s.Contains(3) {
		t.Error()
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error()
	}
}
