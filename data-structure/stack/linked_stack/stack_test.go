package array_stack

import (
	"fmt"
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

	if s.Len() != 3 {
		t.Error()
	}

	a := s.Pop()
	if a != 3 {
		fmt.Println(a)
		t.Error()
	}

	b := s.Peek()
	if b != 2 {
		fmt.Println(b)
		t.Error()
	}

}
