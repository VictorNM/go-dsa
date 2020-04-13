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

// TODO: extract interface
// TODO: implement using linkedList

func TestNew(t *testing.T) {
	s := New(DefaultCapacity)

	assert.IntEqual(t, 0, s.Size())
	assert.True(t, s.IsEmpty())
	assert.IntEqual(t, DefaultCapacity, s.Cap())

	s2 := New(-1)
	assert.IntEqual(t, DefaultCapacity, s2.Cap())
}

func TestPush(t *testing.T) {
	s := New(1)

	ok := s.Push(1)
	assert.True(t, ok)
	assert.IntEqual(t, 1, s.data[s.top])

	ok = s.Push(1)
	assert.False(t, ok)
}

func TestPop(t *testing.T) {
	s := New(1)

	_, ok := s.Pop()
	assert.False(t, ok)

	s.Push(1)
	e, ok := s.Pop()
	assert.True(t, ok)
	assert.IntEqual(t, 1, e)

	_, ok = s.Pop()
	assert.False(t, ok)
}

func TestPeak(t *testing.T) {
	s := New(1)

	_, ok := s.Peak()
	assert.False(t, ok)

	s.Push(10)
	e, ok := s.Peak()
	assert.True(t, ok)
	assert.IntEqual(t, 10, e)
}

func TestIsFull(t *testing.T) {
	s := New(1)
	assert.False(t, s.IsFull())

	s.Push(1)
	assert.True(t, s.IsFull())
}

func TestClear(t *testing.T) {
	s := New(5)

	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)

	s.Clear()
	assert.True(t, s.IsEmpty())
}