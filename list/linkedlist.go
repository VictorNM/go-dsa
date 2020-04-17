package list

type node struct {
	next *node
	e    int
}

type LinkedList struct {
	head *node
	len  int
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *LinkedList) Append(elements ...int) {
	otherL := newFromSlice(elements...)

	if l.len == 0 {
		l.head = otherL.head
		l.len = otherL.len
		return
	}

	p, _ := l.nodeAt(l.len - 1)
	p.next = otherL.head
	l.len += otherL.len
}

func (l *LinkedList) Insert(idx int, e int) (ok bool) {
	if l.outOfRange(idx) && idx != l.Len() {
		return false
	}

	l.len++

	// insert head
	if idx == 0 {
		l.head = &node{
			e:    e,
			next: l.head,
		}
		return true
	}

	// insert middle and tail
	p, _ := l.nodeAt(idx - 1) // 0 <= idx - 1 < l.Len()
	p.next = &node{
		next: p.next,
		e:    e,
	}
	return true
}

func (l *LinkedList) Get(idx int) (e int, ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return 0, false
	}

	return p.e, true
}

func (l *LinkedList) Replace(idx int, e int) (ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return false
	}

	p.e = e
	return true
}

func (l *LinkedList) Remove(idx int) (ok bool) {
	if l.outOfRange(idx) {
		return false
	}

	l.len--

	// remove head
	if idx == 0 {
		l.head = l.head.next
		return true
	}

	p, _ := l.nodeAt(idx - 1)
	p.next = p.next.next
	return true
}

func (l *LinkedList) nodeAt(idx int) (n *node, ok bool) {
	if l.outOfRange(idx) {
		return nil, false
	}

	p := l.head
	for i := 0; i < idx; i++ {
		p = p.next
	}

	return p, true
}

func (l *LinkedList) Search(e int) (idx int, has bool) {
	p := l.head
	idx = 0
	for p != nil {
		if p.e == e {
			return idx, true
		}
		idx++
		p = p.next
	}

	return 0, false
}

func (l *LinkedList) Traverse(f func(e int)) {
	p := l.head
	for p != nil {
		f(p.e)
		p = p.next
	}
}

func (l *LinkedList) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.Len()
}

func NewLinkedList(options ...Option) *LinkedList {
	c := &config{initialSlice: nil}

	for _, o := range options {
		o(c)
	}

	l := newFromSlice(c.initialSlice...)

	return l
}

func newFromSlice(elements ...int) *LinkedList {
	l := &LinkedList{
		len:  0,
		head: nil,
	}

	for i := len(elements) - 1; i >= 0; i-- {
		l.Insert(0, elements[i]) // insert from head is faster
	}

	return l
}
