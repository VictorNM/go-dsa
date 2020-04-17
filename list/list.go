package list

type List interface {
	// Append elements at the end of the list
	Append(elements ...int)

	// Insert the value e at the given index
	// If idx is negative or bigger than the current len of the list, return false and not insert anything
	// If idx equal to the current len of the list, append e at the end of the list
	Insert(idx int, e int) (ok bool)

	// Remove the value at the given index, shift all subsequence elements to the left
	// If idx is negative, bigger of equal to current len of the list, return false and not insert anything
	Remove(idx int) (ok bool)

	// Search return the index of the first element equal to e
	// If not found, return has = false
	Search(e int) (idx int, has bool)

	// Len return the current number of elements in list
	Len() int

	// Get value at the given index
	// If index is negative, bigger or equal to len of the list, return ok = false
	Get(idx int) (e int, ok bool)

	// Replace the element at the given index with new value
	Replace(idx int, e int) (ok bool)

	// Traverse through the list and apply function f
	// Traverse does not change the elements of the list
	Traverse(f func(e int))

	// IsEmpty check if the list is empty
	IsEmpty() bool
}

type config struct {
	capacity     int
	initialSlice []int
}

type Option func(c *config)

// WithInitialCapacity change the default initial capacity of the list
// Only have affect for ArrayList
func WithInitialCapacity(capacity int) Option {
	return func(c *config) {
		c.capacity = capacity
	}
}

// WithInitialSlice will construct a list with the given slice
func WithInitialSlice(slice []int) Option {
	return func(c *config) {
		c.initialSlice = slice
	}
}
