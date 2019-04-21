package b_tree

import (
	"fmt"
	"testing"
)

func comparator(c1 interface{}, c2 interface{}) int {
	a := c1.(int)
	b := c2.(int)
	if a > b {
		return 1

	} else if a < b {
		return -1
	}
	return 0
}

func TestBTree(t *testing.T) {

	tree := New(5, comparator)
	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")

	fmt.Println(tree)

	tests := [][]interface{}{
		{0, nil, false},
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
	}

	for _, test := range tests {
		if value, found := tree.Get(test[0]); value != test[1] || found != test[2] {
			t.Errorf("Got %v,%v expected %v,%v", value, found, test[1], test[2])
		}
	}

	tree.Remove(3)

	_, found := tree.Get(3)
	if found {
		t.Error("Search error")
	}

	if tree.size != 6 {
		t.Error("Error size")
	}

}
