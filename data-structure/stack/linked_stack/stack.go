package array_stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	elements *list.List
	len      int
	lock     sync.Mutex
}

func New() *Stack {
	return &Stack{
		elements: list.New(),
		len:      0,
	}
}

func (s *Stack) Len() int {

	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len

}

func (s *Stack) isEmpty() bool {

	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0

}

func (s *Stack) Pop() (el interface{}) {

	s.lock.Lock()
	defer s.lock.Unlock()
	el = s.elements.Front().Value
	s.elements.Remove(s.elements.Front())
	s.len--
	return el

}

func (s *Stack) Push(el interface{}) {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements.PushFront(el)
	s.len++

}

func (s *Stack) Peek() interface{} {

	s.lock.Lock()
	defer s.lock.Unlock()
	return s.elements.Front().Value

}
