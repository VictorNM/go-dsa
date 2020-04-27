package list

type dllNode struct {
	next *dllNode
	prev *dllNode

	e int
}

// DLL mean Doubly Linked List
type DLL struct {
	head *dllNode
	tail *dllNode

	len int
}

func NewDLL(options ...Option) *DLL {
	config := &config{}

	for _, o := range options {
		o(config)
	}

	l := &DLL{}
	l.Append(config.initialSlice...)

	return l
}

func (l *DLL) Append(elements ...int) {
	for _, e := range elements {
		l.Insert(l.Len(), e)
	}
}

func (l *DLL) Insert(idx int, e int) (ok bool) {
	if idx < 0 || idx > l.Len() {
		return false
	}

	n := &dllNode{e: e}

	if l.IsEmpty() {
		l.head = n
		l.tail = n
		l.len++
		return true
	}

	if idx == 0 {
		n.next = l.head
		l.head.prev = n
		l.head = n
	} else if idx == l.len {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	} else {
		p, _ := l.nodeAt(idx - 1)
		n.next = p.next
		p.next = n
		n.prev = p
	}

	l.len++
	return true
}

func (l *DLL) Remove(idx int) (ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return false
	}

	if p == l.head {
		l.head = p.next
	} else {
		p.prev.next = p.next
	}

	if p == l.tail {
		l.tail = p.prev
	} else {
		p.next.prev = p.prev
	}

	l.len--
	return true
}

func (l *DLL) Search(e int) (idx int, has bool) {
	idx = 0
	for p := l.head; p != nil; p = p.next {
		if p.e == e {
			return idx, true
		}
		idx++
	}

	return 0, false
}

func (l *DLL) Len() int {
	return l.len
}

func (l *DLL) Get(idx int) (e int, ok bool) {
	p, ok := l.nodeAt(idx)
	if ok {
		return p.e, true
	}

	return 0, false
}

func (l *DLL) Replace(idx int, e int) (ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return false
	}

	p.e = e

	return true
}

func (l *DLL) Traverse(f func(e int)) {
	for p := l.head; p != nil; p = p.next {
		f(p.e)
	}
}

func (l *DLL) IsEmpty() bool {
	return l.Len() == 0
}

func (l *DLL) nodeAt(idx int) (n *dllNode, ok bool) {
	if l.outOfRange(idx) {
		return nil, false
	}

	// if idx is in the first half of the list
	// traverse from head
	if idx <= l.Len() / 2 {
		p := l.head
		for i := 0; i < idx; i++ {
			p = p.next
		}

		return p, true
	}

	// else traverse from tail
	p := l.tail
	for i := l.Len() - 1; i > idx; i-- {
		p = p.prev
	}

	return p, true
}

func (l *DLL) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.Len()
}