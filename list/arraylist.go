package list

const DefaultInitialCapacity = 10

type ArrayList struct {
	len      int
	elements []int
}

// Append the value e at the end of the list
func (l *ArrayList) Append(elements ...int) {
	l.ensureCapacity(l.Len() + len(elements))
	copy(l.elements[l.Len():], elements)
	l.len += len(elements)
}

// Insert the value e at index idx, shift all subsequence elements to the right
// If idx is negative or bigger than the current len of the list, return false and not insert anything
// If idx equal to the current len of the list, append e at the end of the list
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

// Remove the value at the given index, shift all subsequence elements to the left
// If idx is negative, bigger of equal to current len of the list, return false and not insert anything
func (l *ArrayList) Remove(idx int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	copy(l.elements[idx:], l.elements[idx+1:])
	l.len--

	l.pack()

	return true
}

// Search return the index of the first element equal to e
// If not found, return has = false
func (l *ArrayList) Search(e int) (idx int, has bool) {
	for i := 0; i < l.len; i++ {
		if l.elements[i] == e {
			return i, true
		}
	}

	return 0, false
}

// Len return the current number of elements in list
func (l *ArrayList) Len() int {
	return l.len
}

// Get value at the given index
// If index is negative, bigger or equal to len of the list, return ok = false
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

// Traverse through the list and apply function f
// Traverse not change the elements of the list
func (l *ArrayList) Traverse(f func(e int)) {
	for i := 0; i < l.len; i++ {
		f(l.elements[i])
	}
}

// IsEmpty check if the list is empty
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

type config struct {
	capacity     int
	initialSlice []int
}

type Option func(c *config)

// WithInitialCapacity change the default initial capacity of the list
func WithInitialCapacity(capacity int) Option {
	return func(c *config) {
		c.capacity = capacity
	}
}

// WithInitialSlice will construct a list with the given slice
func WithInitialSlice(slice []int) Option {
	return func(c *config) {
		c.initialSlice = slice
	}
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
