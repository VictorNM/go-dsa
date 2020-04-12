package list

const DefaultInitialCapacity = 10

type arrayList struct {
	len      int
	elements []int
}

func (l *arrayList) Append(elements ...int) {
	l.ensureCapacity(l.Len() + len(elements))
	copy(l.elements[l.Len():], elements)
	l.len += len(elements)
}

// Insert shift all subsequence elements to the right then insert
func (l *arrayList) Insert(idx int, e int) (ok bool) {
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
func (l *arrayList) Remove(idx int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	copy(l.elements[idx:], l.elements[idx+1:])
	l.len--

	l.pack()

	return true
}

func (l *arrayList) Search(e int) (idx int, has bool) {
	for i := 0; i < l.len; i++ {
		if l.elements[i] == e {
			return i, true
		}
	}

	return 0, false
}

func (l *arrayList) Len() int {
	return l.len
}

func (l *arrayList) Get(idx int) (e int, ok bool) {
	if l.outOfRange(idx) {
		return 0, false
	}

	return l.elements[idx], true
}

func (l *arrayList) Replace(idx int, e int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	l.elements[idx] = e
	return true
}

func (l *arrayList) Traverse(f func(e int)) {
	for i := 0; i < l.len; i++ {
		f(l.elements[i])
	}
}

func (l *arrayList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *arrayList) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.len
}

func (l *arrayList) ensureCapacity(minCapacity int) {
	if minCapacity <= cap(l.elements) {
		return
	}

	capacity := 3*minCapacity/2 + 1
	newElements := make([]int, capacity)
	copy(newElements, l.elements)

	l.elements = newElements
}

func (l *arrayList) pack() {
	if l.len > cap(l.elements)/2 {
		return
	}

	capacity := 3*l.len/2 + 1
	newElements := make([]int, capacity)
	copy(newElements, l.elements)

	l.elements = newElements
}

// NewArrayList return an array list
func NewArrayList(options ...Option) *arrayList {
	c := &config{
		capacity:     DefaultInitialCapacity,
		initialSlice: nil,
	}

	for _, o := range options {
		o(c)
	}

	l := &arrayList{
		len:      0,
		elements: make([]int, c.capacity),
	}
	l.Append(c.initialSlice...)

	return l
}
