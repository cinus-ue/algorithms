package binary_tree

import (
	"fmt"
)

type Comparator func(c1 interface{}, c2 interface{}) int

type Key interface{}
type Value interface{}

type Node struct {
	left, right, parent *Node
	key                 Key
	value               Value
}

type Tree struct {
	root       *Node
	size       int
	Comparator Comparator
}

func New(comparator Comparator) *Tree {
	return &Tree{
		Comparator: comparator,
	}
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Get(k Key) (value Value, found bool) {
	x := t.root
	for x != nil {
		if t.Comparator(k, x.key) < 0 {
			x = x.left
		} else {
			if k == x.key {
				return x.value, true
			}
			x = x.right
		}
	}
	return nil, false
}

func (t *Tree) Put(k Key, v Value) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if t.Comparator(k, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, key: k, value: v}
	t.size++

	if y == nil {
		t.root = z
		return
	} else if t.Comparator(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
}

func (t *Tree) Remove(k Key) {
	current := t.root
	if current == nil {
		return
	}
	parent := t.root
	var isLeft bool
	for t.Comparator(k, current.key) != 0 {
		parent = current
		result := t.Comparator(k, current.key)

		if result < 0 {
			isLeft = true
			current = current.left
		} else if result > 0 {
			isLeft = false
			current = current.right
		}
		if current == nil {
			return
		}
	}

	t.size--
	if current.left == nil && current.right == nil {
		if current == t.root {
			t.root = nil
		} else if isLeft {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if current.left == nil {
		if current == t.root {
			t.root = current.right
		} else if isLeft {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
	} else if current.right == nil {
		if current == t.root {
			t.root = current.left
		} else if isLeft {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
	} else {
		successor := successor(current)
		if current == t.root {
			t.root = successor
		} else if isLeft {
			parent.left = successor
		} else {
			parent.right = successor
		}
		successor.left = current.left
	}
	return
}

func successor(node *Node) *Node {
	successor := node
	parent := node
	current := node.right
	for current != nil {
		parent = successor
		successor = current
		current = current.left
	}
	if successor != node.right {
		parent.left = successor.right
		successor.right = node.right
	}
	return successor
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) Max() Value {
	if t.IsEmpty() {
		return nil
	}
	return maximum(t.root).value
}

func (t *Tree) Min() Value {
	if t.IsEmpty() {
		return nil
	}
	return minimum(t.root).value
}

func minimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maximum(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (t *Tree) String() string {
	str := "BinaryTree\n"
	if !t.IsEmpty() {
		output(t.root, "", true, &str)
	}
	return str
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v", node.key) + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}
