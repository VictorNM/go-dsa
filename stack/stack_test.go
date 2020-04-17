package stack

import (
	"github.com/victornm/go-dsa/shared/assert"
	"testing"
)

// _TODO: New
// _TODO: Push
// _TODO: Pop
// _TODO: Peak
// _TODO: Size
// _TODO: Empty
// _TODO: Clear
// _TODO: Full

// _TODO: extract interface
// _TODO: implement using linkedList

type NewStackFunc func(capacity int) Stack

var newStackFuncMap = map[string]NewStackFunc{
	"array stack": func(cap int) Stack {
		return NewArrayStack(cap)
	},
	"linked list stack": func(cap int) Stack {
		return NewListStack(cap)
	},
}

func TestNew(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(DefaultCapacity)

			assert.IntEqual(t, 0, s.Size())
			assert.True(t, s.IsEmpty())
			assert.IntEqual(t, DefaultCapacity, s.Cap())

			s2 := newStackFunc(-1)
			assert.IntEqual(t, DefaultCapacity, s2.Cap())
		})
	}
}

func TestPush(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(1)

			ok := s.Push(1)
			assert.True(t, ok)
			e, ok := s.Peak()
			assert.True(t, ok)
			assert.IntEqual(t, 1, e)
			assert.IntEqual(t, 1, s.Size())

			ok = s.Push(1)
			assert.False(t, ok)
		})
	}
}

func TestPop(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(1)

			_, ok := s.Pop()
			assert.False(t, ok)

			s.Push(1)
			e, ok := s.Pop()
			assert.True(t, ok)
			assert.IntEqual(t, 1, e)

			_, ok = s.Pop()
			assert.False(t, ok)
		})
	}
}

func TestPeak(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(1)

			_, ok := s.Peak()
			assert.False(t, ok)

			s.Push(10)
			e, ok := s.Peak()
			assert.True(t, ok)
			assert.IntEqual(t, 10, e)
		})
	}
}

func TestIsFull(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(1)
			assert.False(t, s.IsFull())

			s.Push(1)
			assert.True(t, s.IsFull())
		})
	}
}

func TestClear(t *testing.T) {
	for name, newStackFunc := range newStackFuncMap {
		t.Run(name, func(t *testing.T) {
			s := newStackFunc(5)

			s.Push(1)
			s.Push(1)
			s.Push(1)
			s.Push(1)

			s.Clear()
			assert.True(t, s.IsEmpty())
		})
	}
}
