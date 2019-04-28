package xtypes

import (
	"reflect"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue(10)
	if pq == nil {
		t.Fatal("failed to create queue")
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	q := NewPriorityQueue(1)

	expected := &mockItem{value: "put 1", priority: 10}

	q.Push(expected)

	if q.Len() != 1 {
		t.Fatalf("expected %d, got %d", 1, q.Len())
	}

	v, err := q.Pop()
	if err != nil {
		t.Fatalf("failed to get item: %v", err)
	}

	if !reflect.DeepEqual(expected, v) {
		t.Fatalf("expected %v, got %v", *expected, v.(*mockItem))
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	q := NewPriorityQueue(1)

	q.Push(&mockItem{value: "put 1", priority: 10})

	if q.Len() != 1 {
		t.Fatalf("expected %d, got %d", 1, q.Len())
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	q := NewPriorityQueue(1)

	if v := q.Peek(); v != nil {
		t.Fatal("expected nil, got non-nil")
	}

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	q.Push(v1)
	q.Push(v2)

	if v := q.Peek(); !reflect.DeepEqual(v2, v.(*mockItem)) {
		t.Fatalf("expected %v, got %v", *v2, v.(*mockItem))
	}
}

func TestPriorityQueue_PopPriority(t *testing.T) {
	q := NewPriorityQueue(1)

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	q.Push(v1)
	q.Push(v2)

	v, err := q.Pop()
	if err != nil {
		t.Fatalf("failed to pop queue: %v", err)
	}

	if !reflect.DeepEqual(v2, v.(*mockItem)) {
		t.Fatalf("expected %v, got %v", *v2, v.(*mockItem))
	}
}

func TestPriorityQueue_PopEmpty(t *testing.T) {
	q := NewPriorityQueue(1)

	expected := ErrEmptyQueue

	if _, err := q.Pop(); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestPriorityQueue_PushInvalid(t *testing.T) {
	q := &PriorityQueue{}

	expected := ErrInvalidQueue

	if err := q.Push(&mockItem{value: "put 1", priority: 10}); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestPriorityQueue_Get(t *testing.T) {
	q := NewPriorityQueue(2)

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	q.Push(v1)
	q.Push(v2)

	v, err := q.Get(3)
	if err != nil {
		t.Fatalf("failed to get from queue: %v", err)
	}

	if len(v) != 2 {
		t.Fatalf("expected %d, got %d", 2, len(v))
	}
}

func TestPriorityQueue_GetZero(t *testing.T) {
	q := NewPriorityQueue(2)

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	q.Push(v1)
	q.Push(v2)

	v, err := q.Get(0)
	if err != nil {
		t.Fatalf("failed to get from queue: %v", err)
	}

	if len(v) != 0 {
		t.Fatalf("expected %d, got %d", 2, len(v))
	}
}

func TestPriorityQueue_GetInvalid(t *testing.T) {
	q := &PriorityQueue{}

	expected := ErrInvalidQueue

	if _, err := q.Get(1); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestPriorityQueue_Put(t *testing.T) {
	q := NewPriorityQueue(2)

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	if err := q.Put([]PQItem{v1, v2}...); err != nil {
		t.Fatalf("failed to put to queue: %v", err)
	}

	v, err := q.Get(3)
	if err != nil {
		t.Fatalf("failed to get from queue: %v", err)
	}

	if len(v) != 2 {
		t.Fatalf("expected %d, got %d", 2, len(v))
	}
}

func TestPriorityQueue_PutInvalid(t *testing.T) {
	q := &PriorityQueue{}

	v1 := &mockItem{value: "put 1", priority: 10}
	v2 := &mockItem{value: "put 2", priority: 1}

	expected := ErrInvalidQueue

	if err := q.Put([]PQItem{v1, v2}...); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestPriorityQueue_PutEmpty(t *testing.T) {
	q := NewPriorityQueue(2)

	expected := error(nil)

	if err := q.Put([]PQItem{}...); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	TestPriorityQueue_Push(t)
}

func TestPriorityQueue_Empty(t *testing.T) {
	q := NewPriorityQueue(1)

	if !q.Empty() {
		t.Fatalf("expected empty queue, got non-empty")
	}

	q.Push(&mockItem{value: "put 1", priority: 10})

	if q.Empty() {
		t.Fatalf("expected non-empty queue, got empty")
	}

	q.Pop()

	if !q.Empty() {
		t.Fatalf("expected empty queue, got non-empty")
	}
}
