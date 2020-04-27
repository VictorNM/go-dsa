package queue

import "github.com/victornm/go-dsa/list"

type ListQueue struct {
	l     *list.DLL

	capacity int
}

// TODO: Enqueue(e int) (ok bool)
// TODO: Dequeue() (e int, ok bool)
// TODO: Front() (e int, ok bool)
// TODO: Rear() (e int, ok bool)

func (q *ListQueue) Enqueue(e int) (ok bool) {
	if q.IsFull() {
		return false
	}

	q.l.Append(e)

	return true
}

func (q *ListQueue) Dequeue() (e int, ok bool) {
	e, ok = q.l.Get(0)
	if !ok {
		return 0, false
	}

	q.l.Remove(0)

	return e, true
}

func (q *ListQueue) Front() (e int, ok bool) {
	return q.l.Get(0)
}

func (q *ListQueue) Rear() (e int, ok bool) {
	return q.l.Get(q.l.Len() - 1)
}

func (q *ListQueue) Size() int {
	return q.l.Len()
}

func (q *ListQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ListQueue) Cap() int {
	return q.capacity
}

func (q *ListQueue) IsFull() bool {
	return q.Cap() == q.Size()
}

func NewListQueue(capacity int) *ListQueue {
	return &ListQueue{
		capacity: capacity,
		l: list.NewDLL(),
	}
}
