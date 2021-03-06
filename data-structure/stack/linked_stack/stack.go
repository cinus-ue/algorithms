package array_stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	elements *list.List
	lock     sync.Mutex
}

func New() *Stack {
	return &Stack{
		elements: list.New(),
	}
}

func (s *Stack) Len() int {

	s.lock.Lock()
	defer s.lock.Unlock()

	return s.elements.Len()

}

func (s *Stack) IsEmpty() bool {

	s.lock.Lock()
	defer s.lock.Unlock()

	return s.elements.Len() == 0

}

func (s *Stack) Pop() interface{} {

	s.lock.Lock()
	defer s.lock.Unlock()

	return s.elements.Remove(s.elements.Back())

}

func (s *Stack) Push(el interface{}) {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements.PushBack(el)

}

func (s *Stack) Peek() interface{} {

	s.lock.Lock()
	defer s.lock.Unlock()

	return s.elements.Back().Value

}

func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements.Init()
}
