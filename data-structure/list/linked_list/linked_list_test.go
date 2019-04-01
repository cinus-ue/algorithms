package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {

	l := New()

	if !l.IsEmpty() || l.Length != 0 {
		t.Error()
	}

	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	if l.Get(0) != 3 || l.Get(1) != 2 || l.Get(2) != 1 {
		t.Error()
	}

	k := New()
	k.Append(1)
	k.Append(2)
	k.Append(3)

	if k.Get(0) != 1 || k.Get(1) != 2 || k.Get(2) != 3 {
		t.Error()

	}

	l.Concat(k)
	if l.Len() != 6 {
		t.Error()
	}

	counter := 0
	f := func(node *Node) {
		counter += node.Value.(int)
	}

	l.Each(f)
	if counter != 12 {
		t.Error()
	}

	index1, _ := l.IndexOf(1)
	if index1 != 2 {
		fmt.Println(index1)
		t.Error()
	}

	if l.Contains(4) {
		t.Error()
	}

	l.Remove(2)
	if l.Len() != 5 {
		t.Error()
	}

	counter = 0
	l.Each(f)
	if counter != 11 {
		fmt.Println(counter)
		t.Error()
	}

	l.Clear()
	if l.Len() != 0 {
		t.Error()
	}

}
