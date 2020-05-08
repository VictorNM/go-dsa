// Package sort provides primitives for sorting integer slices in ascending order
package sort

// BubbleSort
func BubbleSort(arr []int) {
	changed := true
	for i := 0; i < len(arr) && changed; i++ {
		changed = false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				changed = true
			}
		}
	}
}

// SelectionSort divides the list into 2 parts: sorted and unsorted.
// For each iteration, choose the min value in the unsorted part
// and swap with the element next to the sorted part.
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
}

// InsertionSort
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// QuickSort
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(arr, l, r)

	quickSort(arr, l, pivot-1)
	quickSort(arr, pivot+1, r)
}

func partition(arr []int, l, r int) int {
	pivot := r

	i := l
	for j := l; j < r; j++ {
		if arr[j] < arr[pivot] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i], arr[pivot] = arr[pivot], arr[i]

	return i
}

// MergeSort
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	m := (l + r) / 2

	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)

	merge(arr, l, m+1, r)
}

func merge(arr []int, l, m, r int) {
	sorted := make([]int, r-l+1)

	i, j := l, m
	cur := 0
	for i < m && j <= r {
		if arr[i] < arr[j] {
			sorted[cur] = arr[i]
			i++
		} else {
			sorted[cur] = arr[j]
			j++
		}

		cur++
	}

	// copy remain elements
	if i < m {
		copy(sorted[cur:], arr[i:m])
	} else {
		copy(sorted[cur:], arr[j:r+1])
	}

	copy(arr[l:r+1], sorted)
}
