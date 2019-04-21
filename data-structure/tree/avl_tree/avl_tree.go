package avl_tree

import "fmt"

type Comparator func(c1 interface{}, c2 interface{}) int

type Key interface{}
type Value interface{}

type Tree struct {
	root       *Node
	size       int
	Comparator Comparator
}

type Node struct {
	key      Key
	value    Value
	parent   *Node
	children [2]*Node
	b        int8
}

func New(comparator Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

func (t *Tree) Put(key Key, value Value) {
	t.put(key, value, nil, &t.root)
}

func (t *Tree) Get(key Key) (value Value, found bool) {
	n := t.root
	for n != nil {
		cmp := t.Comparator(key, n.key)
		switch {
		case cmp == 0:
			return n.value, true
		case cmp < 0:
			n = n.children[0]
		case cmp > 0:
			n = n.children[1]
		}
	}
	return nil, false
}

func (t *Tree) Remove(key Key) {
	t.remove(key, &t.root)
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Left() *Node {
	return t.bottom(0)
}

func (t *Tree) Right() *Node {
	return t.bottom(1)
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) put(key Key, value Value, p *Node, qp **Node) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node{key: key, value: value, parent: p}
		return true
	}

	c := t.Comparator(key, q.key)
	if c == 0 {
		q.key = key
		q.value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	var fix bool
	fix = t.put(key, value, q, &q.children[a])
	if fix {
		return putFix(int8(c), qp)
	}
	return false
}

func (t *Tree) remove(key Key, qp **Node) bool {
	q := *qp
	if q == nil {
		return false
	}

	c := t.Comparator(key, q.key)
	if c == 0 {
		t.size--
		if q.children[1] == nil {
			if q.children[0] != nil {
				q.children[0].parent = q.parent
			}
			*qp = q.children[0]
			return true
		}
		fix := removeMin(&q.children[1], &q.key, &q.value)
		if fix {
			return removeFix(-1, qp)
		}
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := t.remove(key, &q.children[a])

	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin(qp **Node, minKey *Key, minVal *Value) bool {
	q := *qp
	if q.children[0] == nil {
		*minKey = q.key
		*minVal = q.value
		if q.children[1] != nil {
			q.children[1].parent = q.parent
		}
		*qp = q.children[1]
		return true
	}
	fix := removeMin(&q.children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.children[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func removeFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}

	if s.b == -c {
		s.b = 0
		return true
	}

	a := (c + 1) / 2
	if s.children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

func singlerot(c int8, s *Node) *Node {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.children[a]
	s.children[a] = rotate(-c, s.children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.children[a]
	s.children[a] = r.children[a^1]
	if s.children[a] != nil {
		s.children[a].parent = s
	}
	r.children[a^1] = s
	r.parent = s.parent
	s.parent = r
	return r
}

func (t *Tree) bottom(d int) *Node {
	n := t.root
	if n == nil {
		return nil
	}

	for c := n.children[d]; c != nil; c = n.children[d] {
		n = c
	}
	return n
}

func (n *Node) Prev() *Node {
	return n.walk1(0)
}

func (n *Node) Next() *Node {
	return n.walk1(1)
}

func (n *Node) walk1(a int) *Node {
	if n == nil {
		return nil
	}

	if n.children[a] != nil {
		n = n.children[a]
		for n.children[a^1] != nil {
			n = n.children[a^1]
		}
		return n
	}

	p := n.parent
	for p != nil && p.children[a] == n {
		n = p
		p = p.parent
	}
	return p
}

func (t *Tree) String() string {
	str := "AVLTree\n"
	if !t.IsEmpty() {
		output(t.root, "", true, &str)
	}
	return str
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v", node.key) + "\n"
	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.children[0], newPrefix, true, str)
	}
}
