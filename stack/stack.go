package stack

const (
	DefaultCapacity = 10
)

type Stack struct {
	top  int
	data []int
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(e int) (ok bool) {
	if s.IsFull() {
		return false
	}

	s.top++
	s.data[s.top] = e
	return true
}

func (s *Stack) Pop() (e int, ok bool) {
	if s.IsEmpty() {
		return 0, false
	}

	e = s.data[s.top]
	s.top--

	return e, true
}

func (s *Stack) Peak() (e int, ok bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.data[s.top], true
}

func (s *Stack) IsFull() bool {
	return s.Size() == s.Cap()
}

func (s *Stack) Cap() int {
	return cap(s.data)
}

func (s *Stack) Clear() {
	s.top = -1
}

func New(capacity int) *Stack {
	if capacity < 0 {
		capacity = DefaultCapacity
	}

	return &Stack{
		top:  -1,
		data: make([]int, capacity),
	}
}
