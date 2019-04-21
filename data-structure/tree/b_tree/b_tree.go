package b_tree

import (
	"fmt"
	"strings"
)

type Comparator func(c1 interface{}, c2 interface{}) int

type Key interface{}
type Value interface{}

type Entry struct {
	key   Key
	value Value
}

type Node struct {
	parent   *Node
	entries  []*Entry
	children []*Node
}

type Tree struct {
	root       *Node
	Comparator Comparator
	size       int
	max        int
}

func New(order int, comparator Comparator) *Tree {
	if order < 3 {
		panic("Invalid order, should be at least 3")
	}
	return &Tree{max: order, Comparator: comparator}
}

func (t *Tree) Put(k Key, v Value) {
	entry := &Entry{key: k, value: v}

	if t.root == nil {
		t.root = &Node{entries: []*Entry{entry}, children: []*Node{}}
		t.size++
		return
	}

	if t.insert(t.root, entry) {
		t.size++
	}
}

func (t *Tree) Get(k Key) (value Value, found bool) {
	node, index, found := t.searchRecursively(t.root, k)
	if found {
		return node.entries[index].value, true
	}
	return nil, false
}

func (t *Tree) Remove(k Key) {
	node, index, found := t.searchRecursively(t.root, k)
	if found {
		t.delete(node, index)
		t.size--
	}
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) Height() int {
	return t.root.height()
}

func (t *Tree) Left() *Node {
	return t.left(t.root)
}

func (t *Tree) LeftKey() Key {
	if left := t.Left(); left != nil {
		return left.entries[0].key
	}
	return nil
}

func (t *Tree) LeftValue() Value {
	if left := t.Left(); left != nil {
		return left.entries[0].value
	}
	return nil
}

func (t *Tree) Right() *Node {
	return t.right(t.root)
}

func (t *Tree) RightKey() Key {
	if right := t.Right(); right != nil {
		return right.entries[len(right.entries)-1].key
	}
	return nil
}

func (t *Tree) RightValue() Value {
	if right := t.Right(); right != nil {
		return right.entries[len(right.entries)-1].value
	}
	return nil
}

func (n *Node) height() int {
	height := 0
	for ; n != nil; n = n.children[0] {
		height++
		if len(n.children) == 0 {
			break
		}
	}
	return height
}

func (t *Tree) isLeaf(node *Node) bool {
	return len(node.children) == 0
}

func (t *Tree) isFull(node *Node) bool {
	return len(node.entries) == t.maxentries()
}

func (t *Tree) shouldSplit(node *Node) bool {
	return len(node.entries) > t.maxentries()
}

func (t *Tree) maxChildren() int {
	return t.max
}

func (t *Tree) minChildren() int {
	return (t.max + 1) / 2 // ceil(m/2)
}

func (t *Tree) maxentries() int {
	return t.maxChildren() - 1
}

func (t *Tree) minentries() int {
	return t.minChildren() - 1
}

func (t *Tree) middle() int {
	return (t.max - 1) / 2 // "-1" to favor right nodes to have more keys when splitting
}

func (t *Tree) search(n *Node, k Key) (index int, found bool) {
	low, high := 0, len(n.entries)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := t.Comparator(k, n.entries[mid].key)
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

func (t *Tree) searchRecursively(start *Node, k Key) (node *Node, index int, found bool) {
	if t.IsEmpty() {
		return nil, -1, false
	}
	node = start
	for {
		index, found = t.search(node, k)
		if found {
			return node, index, true
		}
		if t.isLeaf(node) {
			return nil, -1, false
		}
		node = node.children[index]
	}
}

func (t *Tree) insert(node *Node, entry *Entry) (inserted bool) {
	if t.isLeaf(node) {
		return t.insertIntoLeaf(node, entry)
	}
	return t.insertIntoInternal(node, entry)
}

func (t *Tree) insertIntoLeaf(node *Node, entry *Entry) (inserted bool) {
	position, found := t.search(node, entry.key)
	if found {
		node.entries[position] = entry
		return false
	}

	node.entries = append(node.entries, nil)
	copy(node.entries[position+1:], node.entries[position:])
	node.entries[position] = entry
	t.split(node)
	return true
}

func (t *Tree) insertIntoInternal(node *Node, entry *Entry) (inserted bool) {
	position, found := t.search(node, entry.key)
	if found {
		node.entries[position] = entry
		return false
	}
	return t.insert(node.children[position], entry)
}

func (t *Tree) split(node *Node) {
	if !t.shouldSplit(node) {
		return
	}

	if node == t.root {
		t.splitRoot()
		return
	}

	t.splitNonRoot(node)
}

func (t *Tree) splitNonRoot(node *Node) {
	middle := t.middle()
	parent := node.parent

	left := &Node{entries: append([]*Entry(nil), node.entries[:middle]...), parent: parent}
	right := &Node{entries: append([]*Entry(nil), node.entries[middle+1:]...), parent: parent}

	if !t.isLeaf(node) {
		left.children = append([]*Node(nil), node.children[:middle+1]...)
		right.children = append([]*Node(nil), node.children[middle+1:]...)
		setParent(left.children, left)
		setParent(right.children, right)
	}

	position, _ := t.search(parent, node.entries[middle].key)

	parent.entries = append(parent.entries, nil)
	copy(parent.entries[position+1:], parent.entries[position:])
	parent.entries[position] = node.entries[middle]

	parent.children[position] = left

	parent.children = append(parent.children, nil)
	copy(parent.children[position+2:], parent.children[position+1:])
	parent.children[position+1] = right

	t.split(parent)
}

func (t *Tree) splitRoot() {
	middle := t.middle()

	left := &Node{entries: append([]*Entry(nil), t.root.entries[:middle]...)}
	right := &Node{entries: append([]*Entry(nil), t.root.entries[middle+1:]...)}

	if !t.isLeaf(t.root) {
		left.children = append([]*Node(nil), t.root.children[:middle+1]...)
		right.children = append([]*Node(nil), t.root.children[middle+1:]...)
		setParent(left.children, left)
		setParent(right.children, right)
	}

	newRoot := &Node{
		entries:  []*Entry{t.root.entries[middle]},
		children: []*Node{left, right},
	}

	left.parent = newRoot
	right.parent = newRoot
	t.root = newRoot
}

func setParent(nodes []*Node, parent *Node) {
	for _, node := range nodes {
		node.parent = parent
	}
}

func (t *Tree) left(node *Node) *Node {
	if t.IsEmpty() {
		return nil
	}
	current := node
	for {
		if t.isLeaf(current) {
			return current
		}
		current = current.children[0]
	}
}

func (t *Tree) right(node *Node) *Node {
	if t.IsEmpty() {
		return nil
	}
	current := node
	for {
		if t.isLeaf(current) {
			return current
		}
		current = current.children[len(current.children)-1]
	}
}

func (t *Tree) leftSibling(node *Node, key Key) (*Node, int) {
	if node.parent != nil {
		index, _ := t.search(node.parent, key)
		index--
		if index >= 0 && index < len(node.parent.children) {
			return node.parent.children[index], index
		}
	}
	return nil, -1
}

func (t *Tree) rightSibling(node *Node, key Key) (*Node, int) {
	if node.parent != nil {
		index, _ := t.search(node.parent, key)
		index++
		if index < len(node.parent.children) {
			return node.parent.children[index], index
		}
	}
	return nil, -1
}

func (t *Tree) delete(node *Node, index int) {
	if t.isLeaf(node) {
		deletedKey := node.entries[index].key
		t.deleteEntry(node, index)
		t.rebalance(node, deletedKey)
		if len(t.root.entries) == 0 {
			t.root = nil
		}
		return
	}

	leftLargestNode := t.right(node.children[index])
	leftLargestEntryIndex := len(leftLargestNode.entries) - 1
	node.entries[index] = leftLargestNode.entries[leftLargestEntryIndex]
	deletedKey := leftLargestNode.entries[leftLargestEntryIndex].key
	t.deleteEntry(leftLargestNode, leftLargestEntryIndex)
	t.rebalance(leftLargestNode, deletedKey)
}

func (t *Tree) rebalance(node *Node, deletedKey Key) {
	if node == nil || len(node.entries) >= t.minentries() {
		return
	}

	leftSibling, leftSiblingIndex := t.leftSibling(node, deletedKey)
	if leftSibling != nil && len(leftSibling.entries) > t.minentries() {
		node.entries = append([]*Entry{node.parent.entries[leftSiblingIndex]}, node.entries...)
		node.parent.entries[leftSiblingIndex] = leftSibling.entries[len(leftSibling.entries)-1]
		t.deleteEntry(leftSibling, len(leftSibling.entries)-1)
		if !t.isLeaf(leftSibling) {
			leftSiblingRightMostChild := leftSibling.children[len(leftSibling.children)-1]
			leftSiblingRightMostChild.parent = node
			node.children = append([]*Node{leftSiblingRightMostChild}, node.children...)
			t.deleteChild(leftSibling, len(leftSibling.children)-1)
		}
		return
	}

	rightSibling, rightSiblingIndex := t.rightSibling(node, deletedKey)
	if rightSibling != nil && len(rightSibling.entries) > t.minentries() {

		node.entries = append(node.entries, node.parent.entries[rightSiblingIndex-1])
		node.parent.entries[rightSiblingIndex-1] = rightSibling.entries[0]
		t.deleteEntry(rightSibling, 0)
		if !t.isLeaf(rightSibling) {
			rightSiblingLeftMostChild := rightSibling.children[0]
			rightSiblingLeftMostChild.parent = node
			node.children = append(node.children, rightSiblingLeftMostChild)
			t.deleteChild(rightSibling, 0)
		}
		return
	}

	if rightSibling != nil {

		node.entries = append(node.entries, node.parent.entries[rightSiblingIndex-1])
		node.entries = append(node.entries, rightSibling.entries...)
		deletedKey = node.parent.entries[rightSiblingIndex-1].key
		t.deleteEntry(node.parent, rightSiblingIndex-1)
		t.appendChildren(node.parent.children[rightSiblingIndex], node)
		t.deleteChild(node.parent, rightSiblingIndex)
	} else if leftSibling != nil {

		entries := append([]*Entry(nil), leftSibling.entries...)
		entries = append(entries, node.parent.entries[leftSiblingIndex])
		node.entries = append(entries, node.entries...)
		deletedKey = node.parent.entries[leftSiblingIndex].key
		t.deleteEntry(node.parent, leftSiblingIndex)
		t.prependChildren(node.parent.children[leftSiblingIndex], node)
		t.deleteChild(node.parent, leftSiblingIndex)
	}

	if node.parent == t.root && len(t.root.entries) == 0 {
		t.root = node
		node.parent = nil
		return
	}

	t.rebalance(node.parent, deletedKey)
}

func (t *Tree) prependChildren(fromNode *Node, toNode *Node) {
	children := append([]*Node(nil), fromNode.children...)
	toNode.children = append(children, toNode.children...)
	setParent(fromNode.children, toNode)
}

func (t *Tree) appendChildren(fromNode *Node, toNode *Node) {
	toNode.children = append(toNode.children, fromNode.children...)
	setParent(fromNode.children, toNode)
}

func (t *Tree) deleteEntry(node *Node, index int) {
	copy(node.entries[index:], node.entries[index+1:])
	node.entries[len(node.entries)-1] = nil
	node.entries = node.entries[:len(node.entries)-1]
}

func (t *Tree) deleteChild(node *Node, index int) {
	if index >= len(node.children) {
		return
	}
	copy(node.children[index:], node.children[index+1:])
	node.children[len(node.children)-1] = nil
	node.children = node.children[:len(node.children)-1]
}

func (t *Tree) String() string {
	str := "BTree\n"
	if !t.IsEmpty() {
		output(t.root, 0, &str)
	}
	return str
}

func output(node *Node, level int, str *string) {
	for e := 0; e < len(node.entries)+1; e++ {
		if e < len(node.children) {
			output(node.children[e], level+1, str)
		}
		if e < len(node.entries) {
			*str += strings.Repeat("    ", level)
			*str += fmt.Sprintf("%v", node.entries[e].key) + "\n"
		}
	}
}
