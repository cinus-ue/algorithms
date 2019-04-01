package array_stack

import (
	"github.com/cinus-ue/algorithms-go/data-structure/list/linked_list"
	"sync"
)

type Stack struct {
	elements *linked_list.List
	len      int
	lock     sync.Mutex
}

func New() *Stack {
	return &Stack{
		elements: linked_list.New(),
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

	el = s.elements.Get(0)
	s.elements.Remove(0)
	s.len--
	return

}

func (s *Stack) Push(el interface{}) {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements.Prepend(el)
	s.len++

}

func (s *Stack) Peek() interface{} {

	s.lock.Lock()
	defer s.lock.Unlock()
	return s.elements.Get(0)

}
