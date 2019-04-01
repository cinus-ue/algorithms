package array_list

import "errors"

const (
	growth = float32(2.0)
	shrink = float32(0.25)
)

type List struct {
	elements []interface{}
	Length   int
}

func New() *List {
	return &List{}
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Length == 0
}

func (l *List) Add(values ...interface{}) {
	l.grow(len(values))
	for _, value := range values {
		l.elements[l.Length] = value
		l.Length++
	}
}

func (l *List) Remove(index int) interface{} {
	if !l.rangeCheck(index) {
		return nil
	}
	oldValue := l.elements[index]
	l.elements[index] = nil
	//copy(l.elements[index:], l.elements[index+1:l.Length])
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	l.Length--
	l.shrink()
	return oldValue
}

func (l *List) Get(index int) interface{} {
	if !l.rangeCheck(index) {
		return nil
	}
	return l.elements[index]
}

func (l *List) Set(index int, value interface{}) interface{} {
	if !l.rangeCheck(index) {
		return nil
	}
	oldValue := l.elements[index]
	l.elements[index] = value
	return oldValue
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
	for index <= l.Length {
		if l.elements[index] == o && found == -1 {
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
	l.elements = append([]interface{}{})
}

func (l *List) rangeCheck(index int) bool {
	return index >= 0 && index < l.Length
}

func (l *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *List) grow(n int) {
	currentCapacity := cap(l.elements)
	if l.Length+n >= currentCapacity {
		newCapacity := int(growth * float32(currentCapacity+n))
		l.resize(newCapacity)
	}
}

func (l *List) shrink() {
	if shrink == 0.0 {
		return
	}
	currentCapacity := cap(l.elements)
	if l.Length <= int(float32(currentCapacity)*shrink) {
		l.resize(l.Length)
	}

}
