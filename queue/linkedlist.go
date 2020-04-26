package queue

type node struct {
	prev *node
	next *node

	e int
}

var _ Queue = &ListQueue{}

type ListQueue struct {
	front *node
	rear  *node

	capacity int
	size     int
}

// TODO: Enqueue(e int) (ok bool)
// TODO: Dequeue() (e int, ok bool)
// TODO: Front() (e int, ok bool)
// TODO: Rear() (e int, ok bool)

func (q *ListQueue) Enqueue(e int) (ok bool) {
	if q.IsFull() {
		return false
	}

	n := &node{
		prev: q.rear,
		next: nil,
		e:    e,
	}

	if q.IsEmpty() {
		q.front = n
		q.rear = n
		q.size++

		return true
	}

	q.rear.next = n
	q.rear = n
	q.size++

	return true
}

func (q *ListQueue) Dequeue() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	e = q.front.e
	q.size--

	q.front = q.front.next

	return e, true
}

func (q *ListQueue) Front() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.front.e, true
}

func (q *ListQueue) Rear() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.rear.e, true
}

func (q *ListQueue) Size() int {
	return q.size
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
		size:     0,
	}
}
