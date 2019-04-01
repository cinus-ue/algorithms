package array_queue

import "sync"

type Queue struct {
	elements []interface{}
	len      int
	lock     sync.Mutex
}

func New() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
		len:      0,
	}
}

func (q *Queue) Len() int {

	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len

}

func (q *Queue) IsEmpty() bool {

	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0

}

func (q *Queue) Pop() (el interface{}) {

	q.lock.Lock()
	defer q.lock.Unlock()
	el, q.elements = q.elements[0], q.elements[1:]
	q.len--
	return

}

func (q *Queue) Push(el interface{}) {

	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = append(q.elements, el)
	q.len++
	return
}

func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.elements[0]

}

func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = append([]interface{}{})
	q.len = 0
}
