package xtypes

import (
	"container/heap"
	"sync"
)

// PQItem defines contract which an item of a queue must implement.
type PQItem interface {
	Priority() int
	Index() int
	SetIndex(idx int)
}

// PriorityQueue is a priority queue implemented using heap.
//
// This is clean and simple thread safe implementation with no magic.
// PriorityQueue MUST be created using constructor.
type PriorityQueue struct {
	mu    sync.Mutex // Protects items.
	items PQItems
}

// NewPriorityQueue creates and inits a new PriorityQueue.
func NewPriorityQueue(hint int) *PriorityQueue {
	pq := &PriorityQueue{
		items: make([]PQItem, 0, hint),
	}

	heap.Init(&pq.items)

	return pq
}

// Get returns up to requested n of items.
func (pq *PriorityQueue) Get(n int) ([]PQItem, error) {
	if n < 1 {
		return []PQItem{}, nil
	}

	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.items == nil {
		return nil, ErrInvalidQueue
	}

	result := make([]PQItem, 0, n)

	for i := 0; i < n; i++ {
		if len(pq.items) == 0 {
			break
		}

		result = append(result, heap.Pop(&pq.items).(PQItem))
	}

	return result, nil
}

// Put adds items to queue.
func (pq *PriorityQueue) Put(items ...PQItem) error {
	if len(items) == 0 {
		return nil
	}

	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.items == nil {
		return ErrInvalidQueue
	}

	for _, item := range items {
		heap.Push(&pq.items, item)
	}

	return nil
}

// Pop returns the first element from the queue. If the queue is empty - an error will be returned.
func (pq *PriorityQueue) Pop() (PQItem, error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.items) == 0 {
		return nil, ErrEmptyQueue
	}

	return heap.Pop(&pq.items).(PQItem), nil
}

// Push adds an item to the queue. If the underlying storage is nil - an error will be returned.
func (pq *PriorityQueue) Push(item PQItem) error {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.items == nil {
		return ErrInvalidQueue
	}

	heap.Push(&pq.items, item)

	return nil
}

// Peek returns the first element without modifying the queue.
func (pq *PriorityQueue) Peek() PQItem {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.items) == 0 {
		return nil
	}

	return pq.items[0]
}

// Len returns the len of the queue.
func (pq *PriorityQueue) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return len(pq.items)
}

// Empty returns true if the queue is empty.
func (pq *PriorityQueue) Empty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return len(pq.items) == 0
}

// PQItems represents the queue items.
type PQItems []PQItem

// Pop implements heap.Interface. Returns the first item.
func (pqi *PQItems) Pop() interface{} {
	n := len(*pqi)

	item := (*pqi)[n-1]
	item.SetIndex(-1)

	// Prevent leaks.
	(*pqi)[n-1], *pqi = nil, (*pqi)[0:n-1]

	return item
}

// Push implements heap.Interface. Inserts the x to the queue.
func (pqi *PQItems) Push(x interface{}) {
	n := len(*pqi)

	item := x.(PQItem)
	item.SetIndex(n)

	*pqi = append(*pqi, item)
}

// Len implements heap.Interface. Returns the len of underlying storage.
func (pqi PQItems) Len() int {
	return len(pqi)
}

// Less implements heap.Interface. Returns true if inequality is met.
func (pqi PQItems) Less(i, j int) bool {
	return pqi[i].Priority() < pqi[j].Priority()
}

// Swap implements heap.Interface. Swaps two elements.
func (pqi PQItems) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[i]

	pqi[i].SetIndex(i)
	pqi[j].SetIndex(j)
}
