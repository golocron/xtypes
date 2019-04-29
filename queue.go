package xtypes

import (
	"sync"
)

// Queue is a simple queue based on slice.
//
// This is clean and simple thread safe implementation with no magic.
// Queue MUST be created using constructor.
type Queue struct {
	mu    sync.Mutex
	items QItems
}

// NewQueue creates and inits a new Queue.
func NewQueue(hint int) *Queue {
	q := &Queue{
		items: make(QItems, 0, hint),
	}

	return q
}

// Get returns up to requested n of items.
func (q *Queue) Get(n int) ([]interface{}, error) {
	if n < 1 {
		return []interface{}{}, nil
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.items == nil {
		return nil, ErrInvalidQueue
	}

	result := make([]interface{}, 0, n)

	for i := 0; i < n; i++ {
		if len(q.items) == 0 {
			break
		}

		result = append(result, q.items.Pop())
	}

	return result, nil
}

// Put adds items to queue.
func (q *Queue) Put(items ...interface{}) error {
	if len(items) == 0 {
		return nil
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.items == nil {
		return ErrInvalidQueue
	}

	for _, item := range items {
		q.items.Push(item)
	}

	return nil
}

// Pop returns the first element from the queue. If the queue is empty - an error will be returned.
func (q *Queue) Pop() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, ErrEmptyQueue
	}

	return q.items.Pop(), nil
}

// Push adds an item to the queue. If the underlying storage is nil - an error will be returned.
func (q *Queue) Push(x interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.items == nil {
		return ErrInvalidQueue
	}

	q.items.Push(x)

	return nil
}

// Peek returns the first element without modifying the queue.
func (q *Queue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	return q.items[0]
}

// Len returns the len of the queue.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items)
}

// Empty returns true if the queue is empty.
func (q *Queue) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items) == 0
}

// QItems represents the queue items.
type QItems []interface{}

// Pop returns the first item.
func (qi *QItems) Pop() interface{} {
	n := len(*qi)
	item := (*qi)[0]

	// Prevent leaks.
	(*qi)[0], *qi = nil, (*qi)[1:n]

	return item
}

// Push inserts the x to the queue.
func (qi *QItems) Push(x interface{}) {
	*qi = append(*qi, x)
}
