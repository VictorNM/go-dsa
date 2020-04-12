package list

type List interface {
	Append(elements ...int)
	Insert(idx int, e int) (ok bool)
	Remove(idx int) (ok bool)
	Search(e int) (idx int, has bool)
	Len() int
	Get(idx int) (e int, ok bool)
	Replace(idx int, e int) (ok bool)
	Traverse(f func(e int))
	IsEmpty() bool
}