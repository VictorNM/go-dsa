package stack

const (
	DefaultCapacity = 10
)

type Stack interface {
	// Push an element to the top of stack
	Push(e int) (ok bool)

	// Pop an element from the top of the stack
	Pop() (e int, ok bool)

	// Peak the element on top of the stack
	Peak() (e int, ok bool)

	// Size return the current size of the stack
	Size() int

	// IsEmpty return if the stack is empty
	IsEmpty() bool

	// Cap return the capacity of the stack
	Cap() int

	// IsFull return if the stack is full
	IsFull() bool

	// Clear the stack
	Clear()
}