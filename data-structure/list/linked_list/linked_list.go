package linked_list

import (
	"errors"
)

type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

func New() *List {
	l := new(List)
	l.Length = 0
	return l
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Length == 0
}

func (l *List) Prepend(value interface{}) {

	node := newNode(value)
	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head

	} else {
		formerHead := l.Head
		formerHead.Prev = node
		node.Next = formerHead
		l.Head = node
	}
	l.Length++

}

func (l *List) Append(value interface{}) {

	node := newNode(value)
	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node
		node.Prev = formerTail
		l.Tail = node
	}
	l.Length++

}

func (l *List) Add(value interface{}, index int) error {

	if index > l.Len() {
		return errors.New("index out of range")
	}
	node := newNode(value)
	if l.Len() == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}
	if l.Len()-1 == index {
		l.Append(value)
		return nil
	}
	nextNode, _ := l.getNode(index)
	prevNode := nextNode.Prev
	prevNode.Next = node
	node.Prev = prevNode
	nextNode.Prev = node
	node.Next = nextNode
	l.Length++
	return nil

}

func (l *List) Remove(index int) interface{} {
	if !l.rangeCheck(index) {
		return nil
	}
	node, err := l.getNode(index)
	if err != nil {
		return nil
	}

	next := node.Next
	prev := node.Prev
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}

	if node == l.Head {
		l.Head = node.Next
	}
	if node == l.Tail {
		l.Tail = node.Prev
	}

	l.Length--

	return node.Value
}

func (l *List) Get(index int) (interface{}) {
	node, err := l.getNode(index)
	if err != nil {
		return nil
	}
	return node.Value
}

func (l *List) Contains(o interface{}) bool {
	_, err := l.IndexOf(o)
	if err != nil {
		return false
	}
	return true
}

func (l *List) IndexOf(o interface{}) (int, error) {
	index := 0
	found := -1
	if l.Len() == 0 {
		return found, errors.New("empty list")
	}
	for node := l.Head; node != nil; node = node.Next {
		if node.Value == o && found == -1 {
			found = index
			break
		}
		index++
	}
	if found == -1 {
		return found, errors.New("item not found")
	}
	return found, nil
}

func (l *List) Clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

func (l *List) Concat(k *List) {
	l.Tail.Next, k.Head.Prev = k.Head, l.Tail
	l.Tail = k.Tail
	l.Length += k.Length
}

func (l *List) Each(f func(node *Node)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}

func newNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (l *List) getNode(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("index out of range")
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node, nil
}

func (l *List) rangeCheck(index int) bool {
	return index >= 0 && index < l.Length
}
