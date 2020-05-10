package heap

type Heap interface {
	Push(e int)
	Pop() (e int, ok bool)
	Peak() (e int, ok bool)
	IsEmpty() bool
}

type config struct {
	initialSlice []int
	isMaxHeap    bool
}

type Option func(c *config)

func WithInitialSlice(slice []int) Option {
	return func(c *config) {
		c.initialSlice = slice
	}
}

func UseMaxHeap() Option {
	return func(c *config) {
		c.isMaxHeap = true
	}
}
