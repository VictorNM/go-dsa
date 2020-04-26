package list

type sllNode struct {
	next *sllNode
	e    int
}

// SLL mean Singly Linked List
type SLL struct {
	head *sllNode
	len  int
}

func (l *SLL) Len() int {
	return l.len
}

func (l *SLL) IsEmpty() bool {
	return l.Len() == 0
}

func (l *SLL) Append(elements ...int) {
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

func (l *SLL) Insert(idx int, e int) (ok bool) {
	if l.outOfRange(idx) && idx != l.Len() {
		return false
	}

	l.len++

	// insert head
	if idx == 0 {
		l.head = &sllNode{
			e:    e,
			next: l.head,
		}
		return true
	}

	// insert middle and tail
	p, _ := l.nodeAt(idx - 1) // 0 <= idx - 1 < l.Len()
	p.next = &sllNode{
		next: p.next,
		e:    e,
	}
	return true
}

func (l *SLL) Get(idx int) (e int, ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return 0, false
	}

	return p.e, true
}

func (l *SLL) Replace(idx int, e int) (ok bool) {
	p, ok := l.nodeAt(idx)
	if !ok {
		return false
	}

	p.e = e
	return true
}

func (l *SLL) Remove(idx int) (ok bool) {
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

func (l *SLL) nodeAt(idx int) (n *sllNode, ok bool) {
	if l.outOfRange(idx) {
		return nil, false
	}

	p := l.head
	for i := 0; i < idx; i++ {
		p = p.next
	}

	return p, true
}

func (l *SLL) Search(e int) (idx int, has bool) {
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

func (l *SLL) Traverse(f func(e int)) {
	p := l.head
	for p != nil {
		f(p.e)
		p = p.next
	}
}

func (l *SLL) outOfRange(idx int) bool {
	return idx < 0 || idx >= l.Len()
}

func NewSLL(options ...Option) *SLL {
	c := &config{initialSlice: nil}

	for _, o := range options {
		o(c)
	}

	l := newFromSlice(c.initialSlice...)

	return l
}

func newFromSlice(elements ...int) *SLL {
	l := &SLL{
		len:  0,
		head: nil,
	}

	for i := len(elements) - 1; i >= 0; i-- {
		l.Insert(0, elements[i]) // insert from head is faster
	}

	return l
}
