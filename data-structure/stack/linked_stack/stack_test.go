package array_stack

import (
	"testing"
)

func TestLinkedStack(t *testing.T) {

	s := New()
	if !s.isEmpty() || s.len != 0 || s.Len() != 0 {
		t.Error()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.elements.Get(0) != 3 || s.elements.Get(1) != 2 || s.elements.Get(2) != 1 {
		t.Error()
	}

	if s.Len() != 3 {
		t.Error()
	}

	a := s.Pop()
	if a != 3 {
		t.Error()
	}

	b := s.Peek()
	if b != 2 {
		t.Error()
	}

}
