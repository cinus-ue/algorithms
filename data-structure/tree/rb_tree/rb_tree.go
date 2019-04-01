package rb_tree

import "fmt"

type Color uint

const RED, BLACK Color = 0, 1

type Key interface {
	LessThan(interface{}) bool
}

type Value interface{}

type Node struct {
	left, right, parent *Node
	color               Color
	key                 Key
	value               Value
}

type RBTree struct {
	root *Node
	size int
}

func NewTree() *RBTree {
	return &RBTree{size: 0}
}

func (t *RBTree) Size() int {
	return t.size
}

func (t *RBTree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *RBTree) Search(k Key) *Node {
	return t.findnode(k)
}

func (t *RBTree) IsEmpty() bool {
	if t.root == nil {
		return true
	}
	return false
}

func (t *RBTree) Iterator() *Node {
	return minimum(t.root)
}

func (t *RBTree) Insert(k Key, v Value) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if k.LessThan(x.key) {
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
	} else if z.key.LessThan(y.key) {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)

}

func (t *RBTree) Remove(k Key) {
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

func (t *RBTree) insertFixup(z *Node) {
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

func (t *RBTree) deleteFixup(x, parent *Node) {
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

func (t *RBTree) leftRotate(x *Node) {
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

func (t *RBTree) rightRotate(x *Node) {
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

func (t *RBTree) findnode(k Key) *Node {
	x := t.root
	for x != nil {
		if k.LessThan(x.key) {
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

func (t *RBTree) transplant(u, v *Node) {
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

func (t *RBTree) Preorder() {
	if t.root != nil {
		t.root.preorder()
	}
}

func (n *Node) preorder() {
	fmt.Printf("key:[%v] value:[%s]", n.key, n.value)
	if n.parent == nil {
		fmt.Printf("parent:[nil]")
	} else {
		fmt.Printf("parent:[%v]", n.parent.key)
	}
	if n.color == RED {
		fmt.Print("color:[RED]")
	} else {
		fmt.Print("color:[BLACK]")
	}
	if n.left != nil {
		fmt.Printf("left child:[%v]\n", n.key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("right child:[%v]\n", n.key)
		n.right.preorder()
	}
}
