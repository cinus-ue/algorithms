package array_stack

import "sync"

type Stack struct {
	elements []interface{}
	len      int
	lock     sync.Mutex
}

func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
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

	el, s.elements = s.elements[0], s.elements[1:]
	s.len--
	return

}

func (s *Stack) Push(el interface{}) {

	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := make([]interface{}, 1)
	prepend[0] = el
	s.elements = append(prepend, s.elements...)
	s.len++

}

func (s *Stack) Peek() interface{} {

	s.lock.Lock()
	defer s.lock.Unlock()
	return s.elements[0]

}
