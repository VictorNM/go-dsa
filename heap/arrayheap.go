package heap

type ArrayHeap struct {
	data []int

	isMaxHeap bool
}

func NewArrayHeap(opts ...Option) *ArrayHeap {
	h := &ArrayHeap{}

	c := &config{}
	for _, o := range opts {
		o(c)
	}

	h.isMaxHeap = c.isMaxHeap
	for _, e := range c.initialSlice {
		h.Push(e)
	}

	return h
}

func (h *ArrayHeap) Push(e int) {
	h.data = append(h.data, e)
	h.upheap()
}

func (h *ArrayHeap) Pop() (e int, ok bool) {
	if h.IsEmpty() {
		return 0, false
	}

	e = h.data[0]

	lastIdx := h.size() - 1
	h.data[0], h.data[lastIdx] = h.data[lastIdx], h.data[0]
	h.data = h.data[:lastIdx]

	h.downheap()

	return e, true
}

func (h *ArrayHeap) upheap() {
	cur := len(h.data) - 1

	for cur > 0 {
		par := (cur - 1) / 2
		if h.shouldBeParent(h.data[par], h.data[cur]) {
			return
		}

		h.data[cur], h.data[par] = h.data[par], h.data[cur]
		cur = par
	}
}

func (h *ArrayHeap) downheap() {
	cur := 0

	for {
		l, r := cur*2+1, cur*2+2
		if l >= h.size() {
			break
		}

		if r >= h.size() {
			if h.shouldBeParent(h.data[cur], h.data[l]) {
				break
			}

			h.data[cur], h.data[l] = h.data[l], h.data[cur]
			cur = l
			continue
		}

		par := h.parentIdx(h.data, cur, l, r)
		if par == cur {
			break
		}

		if par == l {
			h.data[cur], h.data[l] = h.data[l], h.data[cur]
			cur = l
		} else {
			h.data[cur], h.data[r] = h.data[r], h.data[cur]
			cur = r
		}
	}
}

// shouldBeParent return true if a should be the parent node of b
func (h *ArrayHeap) shouldBeParent(a, b int) bool {
	if h.isMaxHeap {
		return a >= b
	}

	return a <= b
}

func (h *ArrayHeap) parentIdx(arr []int, i, j, k int) int {
	if h.shouldBeParent(arr[i], arr[j]) && h.shouldBeParent(arr[i], arr[k]) {
		return i
	}

	if h.shouldBeParent(arr[j], arr[k]) {
		return j
	}

	return k
}

func (h *ArrayHeap) Peak() (e int, ok bool) {
	if h.IsEmpty() {
		return 0, false
	}

	return h.data[0], true
}

func (h *ArrayHeap) IsEmpty() bool {
	return h.size() == 0
}

func (h *ArrayHeap) size() int {
	return len(h.data)
}
