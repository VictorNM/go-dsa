package heap_test

import (
	"github.com/victornm/go-dsa/heap"
	"github.com/victornm/go-dsa/shared/assert"
	"testing"
)

// _TODO: Push
// _TODO: Pop
// _TODO: Build from array
// _TODO: support max heap

// _TODO: extract interface
// TODO: create heap using binary tree

type createHeapFunc func(opts ...heap.Option) heap.Heap

var heapCreators = []struct {
	name   string
	create createHeapFunc
}{
	{
		name: "array heap",
		create: func(opts ...heap.Option) heap.Heap {
			return heap.NewArrayHeap(opts...)
		},
	},
}

func TestNew(t *testing.T) {
	for _, creator := range heapCreators {
		t.Run(creator.name, func(t *testing.T) {
			h := creator.create()

			assert.True(t, h.IsEmpty())
		})
	}
}

func TestPeak(t *testing.T) {
	for _, creator := range heapCreators {
		t.Run(creator.name, func(t *testing.T) {
			h := creator.create()

			_, ok := h.Peak()
			assert.False(t, ok)
		})
	}
}

func TestPush(t *testing.T) {
	tests := map[string]struct {
		arr       []int
		wantedTop int
	}{
		"1 element": {
			arr:       []int{1},
			wantedTop: 1,
		},

		"height = 1 unsorted": {
			arr:       []int{3, 2, 1},
			wantedTop: 1,
		},

		"height = 1 sorted": {
			arr:       []int{1, 2, 3},
			wantedTop: 1,
		},

		"height = 2 unsorted": {
			arr:       []int{4, 3, 2, 1},
			wantedTop: 1,
		},

		"height = 2 sorted": {
			arr:       []int{1, 2, 3, 4},
			wantedTop: 1,
		},
	}

	for _, creator := range heapCreators {
		t.Run(creator.name, func(t *testing.T) {
			for name, test := range tests {
				t.Run(name, func(t *testing.T) {
					h := creator.create(heap.WithInitialSlice(test.arr))

					gotTop, _ := h.Peak()
					assert.IntEqual(t, test.wantedTop, gotTop)
				})
			}
		})
	}
}

func TestPop(t *testing.T) {
	t.Run("test ok", func(t *testing.T) {
		h := heap.NewArrayHeap()

		_, ok := h.Pop()
		assert.False(t, ok)

		h.Push(1)

		_, ok = h.Pop()
		assert.True(t, ok)

		_, ok = h.Pop()
		assert.False(t, ok)
	})

	tests := map[string]struct {
		arr []int

		wanted []int
	}{
		"1 element": {
			arr:    []int{1},
			wanted: []int{1},
		},

		"height = 2 sorted": {
			arr:    []int{1, 2, 3, 4, 5},
			wanted: []int{1, 2, 3, 4, 5},
		},

		"height = 2 reversed": {
			arr:    []int{5, 4, 3, 2, 1},
			wanted: []int{1, 2, 3, 4, 5},
		},

		"height = 2 unsorted": {
			arr:    []int{0, 1, 3, 2, 1},
			wanted: []int{0, 1, 1, 2, 3},
		},
	}

	for _, creator := range heapCreators {
		t.Run(creator.name, func(t *testing.T) {
			for name, test := range tests {
				t.Run(name, func(t *testing.T) {
					h := creator.create(heap.WithInitialSlice(test.arr))

					var got []int
					for e, ok := h.Pop(); ok; e, ok = h.Pop() {
						got = append(got, e)
					}

					assert.SliceIntEqual(t, test.wanted, got)
				})
			}
		})
	}
}

func TestMaxHeap(t *testing.T) {
	for _, creator := range heapCreators {
		t.Run(creator.name, func(t *testing.T) {
			h := heap.NewArrayHeap(heap.UseMaxHeap(), heap.WithInitialSlice([]int{1, 2, 3}))

			e, _ := h.Peak()
			assert.IntEqual(t, 3, e)
		})
	}
}
