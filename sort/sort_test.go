package sort_test

import (
	"fmt"
	"github.com/victornm/go-dsa/shared/assert"
	"github.com/victornm/go-dsa/sort"
	"math/rand"
	"testing"
	"time"
)

// _TODO: Bubble sort
// _TODO: Selection sort
// TODO: Heap sort
// _TODO: Insertion sort
// TODO: Shell sort
// _TODO: Quick sort
// _TODO: Merge sort
// TODO: Benchmark
// TODO: Merge sort concurrent

type sortFunc func(arr []int)

var sorters = []struct {
	name string
	sort sortFunc
}{
	{"bubble sort", sort.BubbleSort},
	{"selection sort", sort.SelectionSort},
	{"insertion sort", sort.InsertionSort},
	{"quick sort", sort.QuickSort},
	{"merge sort", sort.MergeSort},
}

func TestSort(t *testing.T) {
	tests := map[string]struct {
		arr    []int
		wanted []int
	}{
		"empty array": {
			arr:    []int{},
			wanted: []int{},
		},

		"1 element": {
			arr:    []int{1},
			wanted: []int{1},
		},

		"2 unsorted elements": {
			arr:    []int{2, 1},
			wanted: []int{1, 2},
		},

		"2 sorted elements": {
			arr:    []int{1, 2},
			wanted: []int{1, 2},
		},

		"10 unsorted elements": {
			arr:    []int{10, 6, 7, 9, 2, 3, 1, 8, 4, 5},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},

		"10 sorted elements": {
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	// log to see the result of each
	//sort.Log = log.New(os.Stdout, "", log.Lmicroseconds)

	for _, sorter := range sorters {
		t.Run(sorter.name, func(t *testing.T) {
			for name, test := range tests {
				t.Run(name, func(t *testing.T) {
					cp := make([]int, len(test.arr))
					copy(cp, test.arr)

					sorter.sort(cp)

					assert.SliceIntEqual(t, test.wanted, cp)
				})
			}
		})
	}
}

var benchTests = []int{
	10,
	100,
	1000,
	10000,
}

// Benchmark_with_10_random_elements/bubble_sort-4                9998074               116 ns/op
// Benchmark_with_10_random_elements/selection_sort-4             7500360               154 ns/op
// Benchmark_with_10_random_elements/insertion_sort-4            13332903                88.5 ns/op
// Benchmark_with_10_random_elements/quick_sort-4                 6520860               190 ns/op
// Benchmark_with_10_random_elements/merge_sort-4                 1687761               685 ns/op
// Benchmark_with_100_random_elements/bubble_sort-4                 86343             12474 ns/op
// Benchmark_with_100_random_elements/selection_sort-4             133328              8893 ns/op
// Benchmark_with_100_random_elements/insertion_sort-4             190401              5956 ns/op
// Benchmark_with_100_random_elements/quick_sort-4                 600564              1981 ns/op
// Benchmark_with_100_random_elements/merge_sort-4                 131839              9276 ns/op
// Benchmark_with_1000_random_elements/bubble_sort-4                  685           1725533 ns/op
// Benchmark_with_1000_random_elements/selection_sort-4              1557            719030 ns/op
// Benchmark_with_1000_random_elements/insertion_sort-4              2352            500425 ns/op
// Benchmark_with_1000_random_elements/quick_sort-4                 20689             57084 ns/op
// Benchmark_with_1000_random_elements/merge_sort-4                  9234            140041 ns/op
// Benchmark_with_10000_random_elements/bubble_sort-4                   6         188336533 ns/op
// Benchmark_with_10000_random_elements/selection_sort-4               19          62211374 ns/op
// Benchmark_with_10000_random_elements/insertion_sort-4               20          51202115 ns/op
// Benchmark_with_10000_random_elements/quick_sort-4                 1303            872602 ns/op
// Benchmark_with_10000_random_elements/merge_sort-4                  648           1800959 ns/op
func BenchmarkSort_RandomArray(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for _, n := range benchTests {
		name := fmt.Sprintf("Benchmark with %d random elements", n)
		b.Run(name, func(b *testing.B) {
			arr := make([]int, n)
			for i := range arr {
				arr[i] = rand.Int()
			}

			bench(b, arr)
		})
	}
}

// Benchmark_with_10_sorted_elements/bubble_sort-4               15999914                73.2 ns/op
// Benchmark_with_10_sorted_elements/selection_sort-4             7843459               148 ns/op
// Benchmark_with_10_sorted_elements/insertion_sort-4            16439661                71.8 ns/op
// Benchmark_with_10_sorted_elements/quick_sort-4                 4669262               268 ns/op
// Benchmark_with_10_sorted_elements/merge_sort-4                 1767248               708 ns/op
// Benchmark_with_100_sorted_elements/bubble_sort-4               3077095               389 ns/op
// Benchmark_with_100_sorted_elements/selection_sort-4             141177              7778 ns/op
// Benchmark_with_100_sorted_elements/insertion_sort-4            2992516               411 ns/op
// Benchmark_with_100_sorted_elements/quick_sort-4                 102564             11378 ns/op
// Benchmark_with_100_sorted_elements/merge_sort-4                 142828              8542 ns/op
// Benchmark_with_1000_sorted_elements/bubble_sort-4               353001              3173 ns/op
// Benchmark_with_1000_sorted_elements/selection_sort-4              1689            643900 ns/op
// Benchmark_with_1000_sorted_elements/insertion_sort-4            352953              3301 ns/op
// Benchmark_with_1000_sorted_elements/quick_sort-4                  1363            876023 ns/op
// Benchmark_with_1000_sorted_elements/merge_sort-4                 12447             96649 ns/op
// Benchmark_with_10000_sorted_elements/bubble_sort-4               34782             32739 ns/op
// Benchmark_with_10000_sorted_elements/selection_sort-4               18          61944400 ns/op
// Benchmark_with_10000_sorted_elements/insertion_sort-4            33706             34000 ns/op
// Benchmark_with_10000_sorted_elements/quick_sort-4                   12          83835133 ns/op
// Benchmark_with_10000_sorted_elements/merge_sort-4                 1008           1146824 ns/op
func BenchmarkSort_SortedArray(b *testing.B) {
	for _, n := range benchTests {
		name := fmt.Sprintf("Benchmark with %d sorted elements", n)
		b.Run(name, func(b *testing.B) {
			arr := make([]int, n)
			for i := range arr {
				arr[i] = i
			}

			bench(b, arr)
		})
	}
}

// Benchmark_with_10_reserved_order_elements/bubble_sort-4              8695474               141 ns/op
// Benchmark_with_10_reserved_order_elements/selection_sort-4           7546201               160 ns/op
// Benchmark_with_10_reserved_order_elements/insertion_sort-4           9449197               116 ns/op
// Benchmark_with_10_reserved_order_elements/quick_sort-4               4762312               246 ns/op
// Benchmark_with_10_reserved_order_elements/merge_sort-4               1764721               702 ns/op
// Benchmark_with_100_reserved_order_elements/bubble_sort-4              115365              9752 ns/op
// Benchmark_with_100_reserved_order_elements/selection_sort-4           123657              8455 ns/op
// Benchmark_with_100_reserved_order_elements/insertion_sort-4           136359              8857 ns/op
// Benchmark_with_100_reserved_order_elements/quick_sort-4               115384             10158 ns/op
// Benchmark_with_100_reserved_order_elements/merge_sort-4               144577              8563 ns/op
// Benchmark_with_1000_reserved_order_elements/bubble_sort-4               1200           1002519 ns/op
// Benchmark_with_1000_reserved_order_elements/selection_sort-4            1818            606723 ns/op
// Benchmark_with_1000_reserved_order_elements/insertion_sort-4            1212            996701 ns/op
// Benchmark_with_1000_reserved_order_elements/quick_sort-4                1538            759427 ns/op
// Benchmark_with_1000_reserved_order_elements/merge_sort-4                9993            100270 ns/op
// Benchmark_with_10000_reserved_order_elements/bubble_sort-4                10         101499970 ns/op
// Benchmark_with_10000_reserved_order_elements/selection_sort-4             19          58526432 ns/op
// Benchmark_with_10000_reserved_order_elements/insertion_sort-4             10         100702070 ns/op
// Benchmark_with_10000_reserved_order_elements/quick_sort-4                 15          72466593 ns/op
// Benchmark_with_10000_reserved_order_elements/merge_sort-4               1034           1148935 ns/op
func BenchmarkSort_ReversedArray(b *testing.B) {
	for _, n := range benchTests {
		name := fmt.Sprintf("Benchmark with %d reserved order elements", n)
		b.Run(name, func(b *testing.B) {
			arr := make([]int, n)
			for i := range arr {
				arr[i] = n-i
			}

			bench(b, arr)
		})
	}
}

func bench(b *testing.B, arr []int) {
	for _, sorter := range sorters {
		b.Run(sorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(arr))
				copy(cp, arr)

				sorter.sort(cp)
			}
		})
	}
}