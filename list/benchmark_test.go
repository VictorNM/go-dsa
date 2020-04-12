package list

import "testing"

// Append: ArrayList faster
// Benchmark_Append_1000/ArrayList-4                  68972             17079 ns/op
// Benchmark_Append_1000/LinkedList-4                  1188            986543 ns/op
func Benchmark_Append_1000(b *testing.B) {
	b.Run("ArrayList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := NewArrayList()

			for i := 0; i < 1000; i++ {
				l.Append(1)
			}
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := NewLinkedList()

			for i := 0; i < 1000; i++ {
				l.Append(1)
			}
		}
	})
}

// Insert head: LinkedList faster
// Benchmark_InsertHead_1000/ArrayList-4               9985            117749 ns/op
// Benchmark_InsertHead_1000/LinkedList-4             25531             47834 ns/op
func Benchmark_InsertHead_1000(b *testing.B) {
	b.Run("ArrayList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := NewArrayList()

			for i := 0; i < 1000; i++ {
				l.Insert(0, 1)
			}
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := NewLinkedList()

			for i := 0; i < 1000; i++ {
				l.Insert(0, 1)
			}
		}
	})
}

// LinkedList Faster
// Benchmark_InsertMiddle_1000/ArrayList-4              100          29599990 ns/op
// Benchmark_InsertMiddle_1000/LinkedList-4            1394            885947 ns/op
func Benchmark_InsertMiddle_1000(b *testing.B) {
	l := NewArrayList(WithInitialSlice(make([]int, 1000)))
	b.Run("ArrayList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Insert(500, 1)
			}
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		l := NewLinkedList(WithInitialSlice(make([]int, 1000)))
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Insert(500, 1)
			}
		}
	})
}

// Random access: ArrayList faster
// Benchmark_Get_1000/ArrayList-4                   1846168               658 ns/op
// Benchmark_Get_1000/LinkedList-4                     1480            806534 ns/op
func Benchmark_Get_1000(b *testing.B) {
	b.Run("ArrayList", func(b *testing.B) {
		l := NewArrayList(WithInitialSlice(make([]int, 1000)))
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Get(i)
			}
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		l := NewLinkedList(WithInitialSlice(make([]int, 1000)))
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Get(i)
			}
		}
	})
}

// Traverse: ArrayList slightly faster
// Benchmark_Traverse_1000/ArrayList-4                  396           2944488 ns/op
// Benchmark_Traverse_1000/LinkedList-4                 288           4267604 ns/op
func Benchmark_Traverse_1000(b *testing.B) {
	b.Run("ArrayList", func(b *testing.B) {
		l := NewArrayList(WithInitialSlice(make([]int, 1000)))
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Traverse(func(e int) {})
			}
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		l := NewLinkedList(WithInitialSlice(make([]int, 1000)))
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1000; i++ {
				l.Traverse(func(e int) {})
			}
		}
	})
}
