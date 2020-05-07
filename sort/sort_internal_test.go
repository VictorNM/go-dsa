package sort

import (
	"github.com/victornm/go-dsa/shared/assert"
	"testing"
)

func TestQuickSort_Partition(t *testing.T) {
	tests := map[string]struct {
		arr []int

		arrWanted   []int
		pivotWanted int
	}{
		"unsorted": {
			arr: []int{5, 1, 2},

			arrWanted:   []int{1, 2, 5},
			pivotWanted: 1,
		},

		"sorted": {
			arr: []int{1, 2, 3},

			arrWanted:   []int{1, 2, 3},
			pivotWanted: 2,
		},

		"sorted reversed": {
			arr: []int{3, 2, 1},

			arrWanted:   []int{1, 2, 3},
			pivotWanted: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cp := make([]int, len(test.arr))
			copy(cp, test.arr)

			pivot := partition(cp, 0, len(cp)-1)

			assert.SliceIntEqual(t, test.arrWanted, cp)
			assert.IntEqual(t, test.pivotWanted, pivot)
		})
	}
}

func TestMergeSort_Merge(t *testing.T) {
	tests := map[string]struct {
		first  []int
		second []int

		wanted []int
	}{
		"2 reversed element": {
			first:  []int{2},
			second: []int{1},

			wanted: []int{1, 2},
		},

		"": {
			first:  []int{1, 3, 5},
			second: []int{2, 4, 6},

			wanted: []int{1, 2, 3, 4, 5, 6},
		},

		"not equal arrays": {
			first:  []int{1, 2, 5, 7},
			second: []int{3, 4, 6},

			wanted: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fLen, sLen := len(test.first), len(test.second)

			cp := make([]int, fLen+sLen)
			copy(cp[:fLen], test.first)
			copy(cp[fLen:], test.second)

			merge(cp, 0, (len(cp)-1)/2+1, len(cp)-1)

			assert.SliceIntEqual(t, test.wanted, cp)
		})
	}
}
