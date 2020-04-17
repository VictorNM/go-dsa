package list

const DefaultInitialCapacity = 10

type ArrayList struct {
	len      int
	elements []int
}

func (l *ArrayList) Append(elements ...int) {
	l.ensureCapacity(l.Len() + len(elements))
	copy(l.elements[l.Len():], elements)
	l.len += len(elements)
}

// Insert shift all subsequence elements to the right then insert
func (l *ArrayList) Insert(idx int, e int) (ok bool) {
	if l.outOfRange(idx) && idx != l.Len() {
		return false
	}

	l.ensureCapacity(l.Len() + 1)

	// in case idx == len; we don't need to copy
	if idx < l.Len() {
		copy(l.elements[idx+1:], l.elements[idx:])
	}
	l.elements[idx] = e
	l.len++

	return true
}

// Remove and shift all subsequence elements to the left
func (l *ArrayList) Remove(idx int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	copy(l.elements[idx:], l.elements[idx+1:])
	l.len--

	l.pack()

	return true
}

func (l *ArrayList) Search(e int) (idx int, has bool) {
	for i := 0; i < l.len; i++ {
		if l.elements[i] == e {
			return i, true
		}
	}

	return 0, false
}

func (l *ArrayList) Len() int {
	return l.len
}

func (l *ArrayList) Get(idx int) (e int, ok bool) {
	if l.outOfRange(idx) {
		return 0, false
	}

	return l.elements[idx], true
}

func (l *ArrayList) Replace(idx int, e int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	l.elements[idx] = e
	return true
}

func (l *ArrayList) Traverse(f func(e int)) {
	for i := 0; i < l.len; i++ {
		f(l.elements[i])
	}
}

func (l *ArrayList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *ArrayList) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.len
}

func (l *ArrayList) ensureCapacity(minCapacity int) {
	if minCapacity <= cap(l.elements) {
		return
	}

	capacity := 3*minCapacity/2 + 1
	newElements := make([]int, capacity)
	copy(newElements, l.elements)

	l.elements = newElements
}

func (l *ArrayList) pack() {
	if l.len > cap(l.elements)/2 {
		return
	}

	capacity := 3*l.len/2 + 1
	newElements := make([]int, capacity)
	copy(newElements, l.elements)

	l.elements = newElements
}

// NewArrayList return an array list
func NewArrayList(options ...Option) *ArrayList {
	c := &config{
		capacity:     DefaultInitialCapacity,
		initialSlice: nil,
	}

	for _, o := range options {
		o(c)
	}

	l := &ArrayList{
		len:      0,
		elements: make([]int, c.capacity),
	}
	l.Append(c.initialSlice...)

	return l
}
