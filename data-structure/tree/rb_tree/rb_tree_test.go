package rb_tree

import (
	"testing"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func TestSearch(t *testing.T) {

	tree := NewTree()

	tree.Insert(key(1), "erf")
	tree.Insert(key(3), "ksd")
	tree.Insert(key(4), "oiu")
	tree.Insert(key(6), "njk")
	tree.Insert(key(5), "pqu")
	tree.Insert(key(2), "lak")

	n := tree.Search(key(4))
	if n.value != "oiu" {
		t.Error("Error value")
	}
	n.value = "kkk"
	if n.value != "kkk" {
		t.Error("Error value modify")
	}
	value := tree.Search(key(5)).value
	if value != "pqu" {
		t.Error("Error value after modifyed other node")
	}
}

func TestIterator(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "erf")
	tree.Insert(key(3), "ksd")
	tree.Insert(key(4), "oiu")
	tree.Insert(key(6), "njk")
	tree.Insert(key(5), "pqu")
	tree.Insert(key(2), "lak")

	it := tree.Iterator()

	for it != nil {
		it = it.Next()
	}

}

func TestRemove(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "erf")
	tree.Insert(key(3), "ksd")
	tree.Insert(key(4), "oiu")
	tree.Insert(key(6), "njk")
	tree.Insert(key(5), "pqu")
	tree.Insert(key(2), "lak")
	for i := 1; i <= 6; i++ {
		tree.Remove(key(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(key(1), "kkk")
	tree.Clear()
	tree.Preorder()
	if tree.Search(key(1)) != nil {
		t.Error("Can't clear")
	}

	tree.Insert(key(4), "piu")
	tree.Insert(key(2), "gfs")
	tree.Insert(key(3), "lki")
	tree.Insert(key(1), "qwe")
	tree.Insert(key(8), "ytr")
	tree.Insert(key(5), "bhg")
	tree.Insert(key(7), "zli")
	tree.Insert(key(9), "exm")
	tree.Remove(key(1))
	tree.Remove(key(2))
}

func TestPreorder(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "kou")
	tree.Insert(key(3), "ihd")
	tree.Insert(key(4), "awe")
	tree.Insert(key(6), "kih")
	tree.Insert(key(5), "mkj")
	tree.Insert(key(2), "xnc")
	if tree.Size() != 6 {
		t.Error("Error size")
	}
	tree.Preorder()
}
