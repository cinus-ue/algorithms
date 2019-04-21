package rb_tree

import (
	"fmt"
	"testing"
)

func comparator(x interface{}, y interface{}) bool {
	if x.(int) < y.(int) {
		return true
	} else {
		return false
	}

}

func TestRBTree(t *testing.T) {

	tree := New(comparator)
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
