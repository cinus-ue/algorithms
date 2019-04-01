package array_queue

import (
	"fmt"

	"testing"
)

func TestArrayQueue(t *testing.T) {

	q := New()

	if !q.IsEmpty() || q.len != 0 || q.Len() != 0 {
		t.Error()
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.elements[0] != 1 || q.elements[1] != 2 || q.elements[2] != 3 {
		fmt.Println(q.elements)
		t.Error()
	}

	if q.Len() != 3 {
		t.Error()
	}

	a := q.Pop()
	if a != 1 || q.Len() != 2 {
		t.Error()
	}

	b := q.Peek()
	if b != 2 {
		t.Error()
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Error()
	}
}
