package list

type node struct {
	next *node
	e    int
}

type linkedList struct {
	head *node
	len  int
}

func (l *linkedList) Len() int {
	return l.len
}

func (l *linkedList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *linkedList) Append(elements ...int) {
	for _, e := range elements {
		l.appendOne(e)
	}
}

func (l *linkedList) Insert(idx int, e int) (ok bool) {
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

func (l *linkedList) appendOne(e int) {
	n := &node{
		next: nil,
		e:    e,
	}

	if l.head == nil {
		l.head = n
	} else {
		p, _ := l.nodeAt(l.Len() - 1) // last node
		p.next = n
	}

	l.len++
}

func (l *linkedList) Get(idx int) (e int, ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return 0, false
	}

	return p.e, true
}

func (l *linkedList) Replace(idx int, e int) (ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return false
	}

	p.e = e
	return true
}

func (l *linkedList) Remove(idx int) (ok bool) {
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

func (l *linkedList) nodeAt(idx int) (n *node, ok bool) {
	if l.outOfRange(idx) {
		return nil, false
	}

	p := l.head
	for i := 0; i < idx; i++ {
		p = p.next
	}

	return p, true
}

func (l *linkedList) Search(e int) (idx int, has bool) {
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

func (l *linkedList) Traverse(f func(e int)) {
	p := l.head
	for p != nil {
		f(p.e)
		p = p.next
	}
}

func (l *linkedList) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.Len()
}

func NewLinkedList(options ...Option) *linkedList {
	c := &config{initialSlice: nil}

	for _, o := range options {
		o(c)
	}

	l := &linkedList{
		len:  0,
		head: nil,
	}

	for i := len(c.initialSlice) - 1; i >= 0; i-- {
		// insert from head is faster
		l.Insert(0, c.initialSlice[i])
	}

	return l
}
