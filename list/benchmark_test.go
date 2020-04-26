package list_test

import (
	. "github.com/victornm/go-dsa/list"
	"testing"
)

// Append: ArrayList >> DLL >> SLL
// Benchmark_Append_1000/ArrayList-4                  62178             19453 ns/op
// Benchmark_Append_1000/SLL-4                         1262            985585 ns/op
// Benchmark_Append_1000/DLL-4                        18837             61156 ns/op
func Benchmark_Append_1000(b *testing.B) {
	for name, createListFunc := range createListFuncMap {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l := createListFunc()

				for i := 0; i < 1000; i++ {
					l.Append(1)
				}
			}
		})
	}
}

// Insert head: SLL > DLL >> ArrayList
// Benchmark_InsertHead_1000/ArrayList-4               9226            117277 ns/op
// Benchmark_InsertHead_1000/SLL-4                    24175             48894 ns/op
// Benchmark_InsertHead_1000/DLL-4                    20100             56766 ns/op
func Benchmark_InsertHead_1000(b *testing.B) {
	for name, createListFunc := range createListFuncMap {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l := createListFunc()

				for i := 0; i < 1000; i++ {
					l.Insert(0, 1)
				}
			}
		})
	}
}

// SLL > DLL >> ArrayList
// Benchmark_InsertMiddle_1000/ArrayList-4              100          29524908 ns/op
// Benchmark_InsertMiddle_1000/SLL-4                   1411            867485 ns/op
// Benchmark_InsertMiddle_1000/DLL-4                   1363            895099 ns/op
func Benchmark_InsertMiddle_1000(b *testing.B) {
	for name, createListFunc := range createListFuncMap {
		l := createListFunc(WithInitialSlice(make([]int, 1000)))
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {

				for i := 0; i < 1000; i++ {
					l.Insert(0, 1)
				}
			}
		})
	}
}

// Random access: ArrayList >> DLL > SLL
// Benchmark_Get_1000/ArrayList-4                   1853538               652 ns/op
// Benchmark_Get_1000/SLL-4                            1578            784874 ns/op
// Benchmark_Get_1000/DLL-4                            3074            376715 ns/op
func Benchmark_Get_1000(b *testing.B) {
	for name, createListFunc := range createListFuncMap {
		l := createListFunc(WithInitialSlice(make([]int, 1000)))
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < 1000; i++ {
					l.Get(i)
				}
			}
		})
	}
}

// Traverse: ArrayList > SLL = DLL
// Benchmark_Traverse_1000/ArrayList-4                  411           3065746 ns/op
// Benchmark_Traverse_1000/SLL-4                        285           4186032 ns/op
// Benchmark_Traverse_1000/DLL-4                        288           4138980 ns/op
func Benchmark_Traverse_1000(b *testing.B) {
	for name, createListFunc := range createListFuncMap {
		l := createListFunc(WithInitialSlice(make([]int, 1000)))
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < 1000; i++ {
					l.Traverse(func(e int) {})
				}
			}
		})
	}
}
