package stack

type arrayStack struct {
	top  int
	data []int
}

func (s *arrayStack) Size() int {
	return s.top + 1
}

func (s *arrayStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *arrayStack) Push(e int) (ok bool) {
	if s.IsFull() {
		return false
	}

	s.top++
	s.data[s.top] = e
	return true
}

func (s *arrayStack) Pop() (e int, ok bool) {
	if s.IsEmpty() {
		return 0, false
	}

	e = s.data[s.top]
	s.top--

	return e, true
}

func (s *arrayStack) Peak() (e int, ok bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.data[s.top], true
}

func (s *arrayStack) IsFull() bool {
	return s.Size() == s.Cap()
}

func (s *arrayStack) Cap() int {
	return cap(s.data)
}

func (s *arrayStack) Clear() {
	s.top = -1
}

func NewArrayStack(capacity int) *arrayStack {
	if capacity < 0 {
		capacity = DefaultCapacity
	}

	return &arrayStack{
		top:  -1,
		data: make([]int, capacity),
	}
}
