package rb_tree

import "fmt"

type Color uint

const RED, BLACK Color = 0, 1

type Comparator func(c1 interface{}, c2 interface{}) bool

type Key interface{}
type Value interface{}

type Node struct {
	left, right, parent *Node
	color               Color
	key                 Key
	value               Value
}

type Tree struct {
	root       *Node
	size       int
	Comparator Comparator
}

func New(comparator Comparator) *Tree {
	return &Tree{size: 0, Comparator: comparator}
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) Get(k Key) (value Value, found bool) {
	node := t.findnode(k)
	if node == nil {
		return nil, false
	}
	return node.value, true
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Iterator() *Node {
	return minimum(t.root)
}

func (t *Tree) Put(k Key, v Value) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if t.Comparator(k, x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, color: RED, key: k, value: v}
	t.size++

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if t.Comparator(z.key, y.key) {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)

}

func (t *Tree) Remove(k Key) {
	z := t.findnode(k)
	if z == nil {
		return
	}

	var x, y, parent *Node
	y = z
	yOriginalColor := y.color
	parent = z.parent
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = minimum(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent == z {
			if x == nil {
				parent = y
			} else {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x, parent)
	}
	t.size--
}

func (t *Tree) insertFixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *Tree) deleteFixup(x, parent *Node) {
	var w *Node

	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.leftRotate(parent)
				w = parent.right
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.right) == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					t.rightRotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				t.leftRotate(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.rightRotate(parent)
				w = parent.left
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.left) == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					t.leftRotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				t.rightRotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *Tree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

func (t *Tree) findnode(k Key) *Node {
	x := t.root
	for x != nil {
		if t.Comparator(k, x.key) {
			x = x.left
		} else {
			if k == x.key {
				return x
			}
			x = x.right
		}
	}
	return nil
}

func (t *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v == nil {
		return
	}
	v.parent = u.parent
}

func (n *Node) Next() *Node {
	return successor(n)
}

func successor(x *Node) *Node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

func getColor(n *Node) Color {
	if n == nil {
		return BLACK
	}
	return n.color
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
	str := "RBTree\n"
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
