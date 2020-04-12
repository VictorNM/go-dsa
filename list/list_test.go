package list_test

import (
	"fmt"
	. "github.com/victornm/go-dsa/list"
	"testing"
)

// List of TODOs here illustrate the to-do list
// in "Test-Driven Development By Example" - Kent Beck

// Basic operations
// _TODO: Construct a list, leaving it empty.
// _TODO: Insert an element.
// _TODO: Insert Tail
// _TODO: Insert Head
// _TODO: Insert at index
// _TODO: Remove an element.
// _TODO: Search an element.
// _TODO: Retrieve an element.
// _TODO: Traverse the list, performing a given operation on each element.

// Extended operations
// _TODO: Empty or not.
// _TODO: Size of the list.
// _TODO: Replace an element with another element.
// _TODO: Append an unordered list to another.

// _TODO: replace len() with l.Len()

// _TODO: extract List interface, change these tests to test through the interface
// TODO: implement LinkedList by passing all the tests here
// TODO: benchmark to show different between LinkedList and ArrayList

func TestNew(t *testing.T) {
	l := NewArrayList()

	if l.Len() != 0 || !l.IsEmpty() {
		t.Error("new list should be empty")
	}

	l.Append(1)
	assertEqual(t, l.Len(), 1)
	assertFalse(t, l.IsEmpty())
}

func TestNewWithInitialSlice(t *testing.T) {
	tests := map[string]struct {
		capacity int
		slice    []int
	}{
		"cap > len(slice)": {
			capacity: 5,
			slice:    []int{1, 2, 3},
		},

		"cap < len(slice)": {
			capacity: 5,
			slice:    []int{1, 2, 3, 4, 5, 6, 7},
		},

		"cap == len(slice)": {
			capacity: 5,
			slice:    []int{1, 2, 3, 4, 5},
		},

		"nil slice": {
			capacity: 5,
			slice:    nil,
		},

		"empty slice": {
			capacity: 5,
			slice:    []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList(WithInitialCapacity(test.capacity), WithInitialSlice(test.slice))

			assertSliceEqual(t, test.slice, toSlice(l))
		})
	}
}

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		f      func(l List)
		wanted []int
	}{
		"append 1 time": {
			func(l List) {
				l.Append(1)
			},
			[]int{1},
		},

		"append 2 times": {
			func(l List) {
				l.Append(1)
				l.Append(2)
			},
			[]int{1, 2},
		},

		"append multiple elements": {
			f: func(l List) {
				l.Append(1, 2, 3)
			},
			wanted: []int{1, 2, 3},
		},

		"append slice": {
			f: func(l List) {
				l.Append([]int{1, 2, 3}...)
			},
			wanted: []int{1, 2, 3},
		},

		"append nothing": {
			f: func(l List) {
				l.Append(1)
				l.Append()
			},

			wanted: []int{1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList()
			test.f(l)
			assertSliceEqual(t, test.wanted, toSlice(l))
		})
	}
}

func TestInsert(t *testing.T) {
	tests := map[string]struct {
		f      func(l List)
		wanted []int // list state after perform f
	}{
		"insert to empty list at index 0": {
			f: func(l List) {
				l.Insert(0, 1)
			},
			wanted: []int{1},
		},

		"insert at negative index should not successfully": {
			f: func(l List) {
				l.Insert(-1, 1)
			},
			wanted: nil,
		},

		"insert out of range should not successfully": {
			f: func(l List) {
				l.Insert(1, 2)
			},
			wanted: nil,
		},

		"insert many times": {
			f: func(l List) {
				l.Insert(0, 2)
				l.Insert(0, 1)
				l.Insert(1, 3)
				l.Insert(2, 4)
				l.Insert(4, 1)
			},
			wanted: []int{1, 3, 4, 2, 1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList()
			test.f(l)
			assertSliceEqual(t, test.wanted, toSlice(l))
		})
	}
}

func TestReplace(t *testing.T) {
	tests := map[string]struct {
		f      func(l List)
		wanted []int // list state after perform f
	}{
		"replace empty list": {
			f: func(l List) {
				l.Replace(0, 1)
			},
			wanted: nil,
		},

		"replace list with 1 element": {
			f: func(l List) {
				l.Insert(0, 1)
				l.Replace(0, 2)
			},
			wanted: []int{2},
		},

		"replace negative index": {
			f: func(l List) {
				l.Insert(0, 1)
				l.Replace(-1, 1)
			},
			wanted: []int{1},
		},

		"replace out of range index": {
			f: func(l List) {
				l.Insert(0, 1)
				l.Replace(1, 1)
			},
			wanted: []int{1},
		},

		"replace at middle element": {
			f: func(l List) {
				l.Insert(0, 0)
				l.Insert(1, 1)
				l.Insert(2, 2)
				l.Replace(1, 3)
			},
			wanted: []int{0, 3, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList()
			test.f(l)
			assertSliceEqual(t, test.wanted, toSlice(l))
		})
	}
}

func TestRemove(t *testing.T) {
	tests := map[string]struct {
		f      func(l List)
		wanted []int // list state after perform f
	}{
		"remove empty list": {
			f: func(l List) {
				l.Remove(0)
			},
			wanted: nil,
		},

		"append once then remove": {
			f: func(l List) {
				l.Append(1)
				l.Remove(0)
			},
			wanted: nil,
		},

		"append once then remove out of range": {
			f: func(l List) {
				l.Append(0)
				l.Remove(1)
			},
			wanted: []int{0},
		},

		"append twice then remove front": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Remove(0)
			},
			wanted: []int{2},
		},

		"append twice then remove back": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Remove(1)
			},
			wanted: []int{1},
		},

		"append 4 times then remove 3 times": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
				l.Remove(0)
				l.Remove(0)
				l.Remove(0)
			},
			wanted: []int{4},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList(WithInitialCapacity(1))
			test.f(l)
			assertSliceEqual(t, test.wanted, toSlice(l))
		})
	}
}

func TestSearch(t *testing.T) {
	tests := map[string]struct {
		f func(l List)

		searchFor int

		wantedIdx int
		wantedHas bool
	}{
		"empty slice": {
			f:         func(l List) {},
			searchFor: 10,
			wantedHas: false,
		},

		"slice has 1 element": {
			f: func(l List) {
				l.Append(1)
			},
			searchFor: 1,
			wantedHas: true,
			wantedIdx: 0,
		},

		"slice = [2] after remove": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Remove(0)
			},
			searchFor: 2,
			wantedHas: true,
			wantedIdx: 0,
		},

		"removed element should not be found": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Remove(1)
			},
			searchFor: 2,
			wantedHas: false,
		},

		"search for element at the begin of slice": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
			},
			searchFor: 1,
			wantedHas: true,
			wantedIdx: 0,
		},

		"search for element at the end of slice": {
			f: func(l List) {
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
			},
			searchFor: 4,
			wantedHas: true,
			wantedIdx: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList()
			test.f(l)

			gotIdx, gotHas := l.Search(test.searchFor)

			if test.wantedHas {
				assertTrue(t, gotHas)
				assertEqual(t, test.wantedIdx, gotIdx)
			} else {
				assertFalse(t, gotHas)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := map[string]struct {
		f func(l List)

		getAt int

		wantedE   int
		wantedHas bool
	}{
		"get empty slice": {
			f:         func(l List) {},
			getAt:     0,
			wantedHas: false,
		},

		"append once then get at 0": {
			f: func(l List) {
				l.Append(1)
			},
			getAt:     0,
			wantedHas: true,
			wantedE:   1,
		},

		"append once then get out of range": {
			f: func(l List) {
				l.Append(1)
			},
			getAt:     1,
			wantedHas: false,
		},

		"get at negative index": {
			f: func(l List) {
				l.Append(1)
			},
			getAt:     -1,
			wantedHas: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewArrayList()
			test.f(l)

			gotE, has := l.Get(test.getAt)

			if test.wantedHas {
				assertTrue(t, has)
				assertEqual(t, test.wantedE, gotE)
			} else {
				assertFalse(t, has)
			}
		})
	}
}

func TestTraverse(t *testing.T) {
	l := NewArrayList()

	wanted := []int{0, 1, 2, 3}

	var got []int

	for _, e := range wanted {
		l.Append(e)
	}

	l.Traverse(func(e int) {
		got = append(got, e)
	})

	assertSliceEqual(t, wanted, got)
}

// toSlice return the slice presentation of List
// return a nil-slice if l.Len() == 0
func toSlice(l List) []int {
	var slice []int

	for i := 0; i < l.Len(); i++ {
		e, _ := l.Get(i)
		slice = append(slice, e)
	}

	return slice
}

func assertTrue(t *testing.T, condition bool, msg ...string) {
	t.Helper()
	assertCondition(t, true, condition, msg...)
}

func assertFalse(t *testing.T, condition bool, msg ...string) {
	assertCondition(t, false, condition, msg...)
}

func assertCondition(t *testing.T, wanted, got bool, msg ...string) {
	t.Helper()
	if wanted != got {
		if len(msg) == 0 {
			msg = []string{fmt.Sprintf("should be %t", wanted)}
		}
		t.Errorf(msg[0])
	}
}

func assertEqual(t *testing.T, wanted, got int, msg ...string) {
	t.Helper()
	if wanted != got {
		if len(msg) == 0 {
			msg = []string{fmt.Sprintf("wanted %d but got %d", wanted, got)}
		}
		t.Errorf(msg[0])
	}
}

// assertSliceEqual
// note: for testing, nil-slice and empty slice are equal
func assertSliceEqual(t *testing.T, wanted, got []int) {
	t.Helper()
	if len(wanted) != len(got) {
		t.Errorf("wanted %v but got %v", wanted, got)
		return
	}

	for i := range wanted {
		assertEqual(t, wanted[i], got[i], fmt.Sprintf("wanted %v but got %v", wanted, got))
	}
}
