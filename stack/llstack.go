package stack

import "github.com/victornm/go-dsa/list"

type linkedListStack struct {
	Stack    // TODO: remove this embedded interface
	capacity int
	l        *list.LinkedList
}

func NewListStack(capacity int) *linkedListStack {
	if capacity < 0 {
		capacity = DefaultCapacity
	}

	return &linkedListStack{
		capacity: capacity,
		l:        list.NewLinkedList(),
	}
}

func (s *linkedListStack) Push(e int) (ok bool) {
	if s.IsFull() {
		return false
	}

	s.l.Insert(0, e)

	return true
}

func (s *linkedListStack) Peak() (e int, ok bool) {
	return s.l.Get(0)
}

func (s *linkedListStack) Pop() (e int, ok bool) {
	e, ok = s.l.Get(0)
	if !ok {
		return 0, false
	}

	s.l.Remove(0)
	return e, ok
}

func (s *linkedListStack) Size() int {
	return s.l.Len()
}

func (s *linkedListStack) IsEmpty() bool {
	return s.l.IsEmpty()
}

func (s *linkedListStack) Cap() int {
	return s.capacity
}

func (s *linkedListStack) IsFull() bool {
	return s.Size() == s.Cap()
}

func (s *linkedListStack) Clear() {
	for ok := true; ok; ok = s.l.Remove(0) {}
}