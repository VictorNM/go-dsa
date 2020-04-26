package queue

type ArrayQueue struct {
	frontIdx, size int
	data           []int
}

func (q *ArrayQueue) Enqueue(e int) (ok bool) {
	if q.IsFull() {
		return false
	}

	q.size++
	q.data[q.rearIdx()] = e

	return true
}

func (q *ArrayQueue) Dequeue() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	e = q.data[q.frontIdx]
	q.frontIdx++
	q.size--

	// if the frontIdx is out of range
	// reset it
	if q.frontIdx >= q.Cap() {
		q.frontIdx = q.frontIdx % q.Cap()
	}

	return e, true
}

func (q *ArrayQueue) Front() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.data[q.frontIdx], true
}

func (q *ArrayQueue) Rear() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.data[q.rearIdx()], true
}

func (q *ArrayQueue) rearIdx() int {
	return (q.frontIdx + q.size - 1) % q.Cap()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) IsFull() bool {
	return q.Size() == q.Cap()
}

func (q *ArrayQueue) Cap() int {
	return cap(q.data)
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		frontIdx: 0,
		size:     0,
		data:     make([]int, capacity),
	}
}
