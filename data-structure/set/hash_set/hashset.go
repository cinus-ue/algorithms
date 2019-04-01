package hash_set

type Set struct {
	elements map[interface{}]struct{}
}

func New(values ...interface{}) *Set {
	set := &Set{elements: make(map[interface{}]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Add(values ...interface{}) {
	for _, value := range values {
		s.elements[value] = struct{}{}
	}
}

func (s *Set) Remove(values ...interface{}) {
	for _, value := range values {
		delete(s.elements, value)
	}
}

func (s *Set) Contains(values ...interface{}) bool {
	for _, value := range values {
		if _, contains := s.elements[value]; !contains {
			return false
		}
	}
	return true
}

func (s *Set) Clear() {
	s.elements = make(map[interface{}]struct{})
}
