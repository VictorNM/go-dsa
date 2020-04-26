package queue

const DefaultCapacity = 10

type Queue interface {
	Enqueue(e int) (ok bool)
	Dequeue() (e int, ok bool)
	Front() (e int, ok bool)
	Rear() (e int, ok bool)
	IsEmpty() bool
	Size() int
	IsFull() bool
	Cap() int
}
