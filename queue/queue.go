package queue

type Queue struct {
	frontIdx, size int
	data           []int
}

func (q *Queue) Enqueue(e int) (ok bool) {
	if q.IsFull() {
		return false
	}

	q.size++
	q.data[q.rearIdx()] = e

	return true
}

func (q *Queue) Dequeue() (e int, ok bool) {
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

func (q *Queue) Front() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.data[q.frontIdx], true
}

func (q *Queue) Rear() (e int, ok bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.data[q.rearIdx()], true
}

func (q *Queue) rearIdx() int {
	return (q.frontIdx + q.size - 1) % q.Cap()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsFull() bool {
	return q.Size() == q.Cap()
}

func (q *Queue) Cap() int {
	return cap(q.data)
}

func New(capacity int) *Queue {
	return &Queue{
		frontIdx: 0,
		size:     0,
		data:     make([]int, capacity),
	}
}
