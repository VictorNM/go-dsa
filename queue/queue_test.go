package queue_test

import (
	. "github.com/victornm/go-dsa/queue"
	"github.com/victornm/go-dsa/shared/assert"
	"testing"
)

// _TODO: New
// _TODO: Enqueue
// _TODO: Dequeue
// _TODO: IsEmpty
// _TODO: Size
// _TODO: IsFull
// _TODO: Cap

func TestNew(t *testing.T) {
	q := New(10)

	assert.True(t, q.IsEmpty())
}

func TestEnqueue(t *testing.T) {
	t.Run("enqueue once", func(t *testing.T) {
		q := New(10)

		ok := q.Enqueue(1)

		assert.True(t, ok)
		assert.IntEqual(t, 1, q.Size())
		assert.False(t, q.IsEmpty())

		front, ok := q.Front()
		assert.True(t, ok)
		assert.IntEqual(t, 1, front)

		rear, ok := q.Rear()
		assert.True(t, ok)
		assert.IntEqual(t, 1, rear)
	})

	t.Run("enqueue to 0 capacity queue", func(t *testing.T) {
		q := New(0)

		ok := q.Enqueue(1)
		assert.False(t, ok)
		assert.IntEqual(t, 0, q.Size())

		_, ok = q.Front()
		assert.False(t, ok)

		_, ok = q.Rear()
		assert.False(t, ok)
	})

	t.Run("enqueue to full queue", func(t *testing.T) {
		q := New(1)

		q.Enqueue(1)
		ok := q.Enqueue(2)
		assert.False(t, ok)

		front, ok := q.Front()
		assert.True(t, ok)
		assert.IntEqual(t, 1, front)

		rear, ok := q.Rear()
		assert.True(t, ok)
		assert.IntEqual(t, 1, rear)
	})

	t.Run("enqueue to full queue cap = 2", func(t *testing.T) {
		q := New(2)

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		front, _ := q.Front()
		assert.IntEqual(t, 1, front)

		rear, _ := q.Rear()
		assert.IntEqual(t, 2, rear)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("dequeue empty", func(t *testing.T) {
		q := New(1)

		_, ok := q.Dequeue()

		assert.False(t, ok)
	})

	t.Run("dequeue once", func(t *testing.T) {
		q := New(1)

		q.Enqueue(1)
		e, ok := q.Dequeue()

		assert.True(t, ok)
		assert.IntEqual(t, 1, e)
		assert.IntEqual(t, 0, q.Size())
	})

	t.Run("en - de - en", func(t *testing.T) {
		q := New(1)

		q.Enqueue(1)
		q.Dequeue()
		q.Enqueue(2)

		front, _ := q.Front()
		assert.IntEqual(t, 2, front)

		rear, _ := q.Rear()
		assert.IntEqual(t, 2, rear)
	})

	t.Run("en - en - de - en", func(t *testing.T) {
		q := New(2)

		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()
		q.Enqueue(3)

		front, _ := q.Front()
		assert.IntEqual(t, 2, front)

		rear, _ := q.Rear()
		assert.IntEqual(t, 3, rear)
	})

	t.Run("en - en - de - en - de - de", func(t *testing.T) {
		q := New(2)

		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()
		q.Enqueue(3)
		q.Dequeue()
		e, ok := q.Dequeue()

		assert.True(t, ok)
		assert.IntEqual(t, 3, e)
		assert.True(t, q.IsEmpty())
	})
}
