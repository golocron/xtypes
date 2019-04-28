package xtypes

import (
	"container/list"
	"reflect"
	"testing"
)

const (
	size = 1024
)

type mockItem struct {
	index int
	priority int
	value string
}

func TestNewQueue(t *testing.T) {
	q := NewQueue(10)
	if q == nil {
		t.Fatal("failed to create queue")
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue(1)

	expected := &mockItem{value: "put 1"}

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

func TestQueue_Push(t *testing.T) {
	q := NewQueue(1)

	q.Push(&mockItem{value: "put 1"})

	if q.Len() != 1 {
		t.Fatalf("expected %d, got %d", 1, q.Len())
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue(1)

	if v := q.Peek(); v != nil {
		t.Fatal("expected nil, got non-nil")
	}

	v1 := &mockItem{value: "put 1"}
	v2 := &mockItem{value: "put 2"}

	q.Push(v1)
	q.Push(v2)

	if v := q.Peek(); !reflect.DeepEqual(v1, v.(*mockItem)) {
		t.Fatalf("expected %v, got %v", *v1, v.(*mockItem))
	}
}

func TestQueue_PopEmpty(t *testing.T) {
	q := NewQueue(1)

	expected := ErrEmptyQueue

	if _, err := q.Pop(); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestQueue_PushInvalid(t *testing.T) {
	q := &Queue{}

	expected := ErrInvalidQueue

	if err := q.Push(&mockItem{value: "put 1"}); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestQueue_Get(t *testing.T) {
	q := NewQueue(2)

	v1 := &mockItem{value: "put 1"}
	v2 := &mockItem{value: "put 2"}

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

func TestQueue_GetZero(t *testing.T) {
	q := NewQueue(2)

	v1 := &mockItem{value: "put 1"}
	v2 := &mockItem{value: "put 2"}

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

func TestQueue_GetInvalid(t *testing.T) {
	q := &Queue{}

	expected := ErrInvalidQueue

	if _, err := q.Get(1); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestQueue_Put(t *testing.T) {
	q := NewQueue(2)

	v1 := &mockItem{value: "put 1"}
	v2 := &mockItem{value: "put 2"}

	if err := q.Put([]interface{}{v1, v2}...); err != nil {
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

func TestQueue_PutInvalid(t *testing.T) {
	q := &Queue{}

	v1 := &mockItem{value: "put 1"}
	v2 := &mockItem{value: "put 2"}

	expected := ErrInvalidQueue

	if err := q.Put([]interface{}{v1, v2}...); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestQueue_PutEmpty(t *testing.T) {
	q := NewQueue(2)

	expected := error(nil)

	if err := q.Put([]interface{}{}...); err != expected {
		t.Fatalf("expected %v, got %v ", expected, err)
	}
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue(1)

	if !q.Empty() {
		t.Fatalf("expected empty queue, got non-empty")
	}

	q.Push(&mockItem{value: "put 1"})

	if q.Empty() {
		t.Fatalf("expected non-empty queue, got empty")
	}

	q.Pop()

	if !q.Empty() {
		t.Fatalf("expected empty queue, got non-empty")
	}
}

// Benchmarks.
func BenchmarkQueuePush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue(0)

		for n := 0; n < size; n++ {
			q.Push(n)
		}
	}
}

func BenchmarkQueuePushHint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue(size)

		for n := 0; n < size; n++ {
			q.Push(n)
		}
	}
}

// BenchmarkListPush tests a queue based on built-in list.
//
// Yes, it looks like it is faster but with a simple exception - it isn't thread safe.
// The implemented queue is thread safe (but this is not the best possible implementation).
func BenchmarkListPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List

		for n := 0; n < size; n++ {
			q.PushBack(n)
		}
	}
}

func BenchmarkChannelPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make(chan interface{}, size)

		for n := 0; n < size; n++ {
			q <- n
		}
		close(q)
	}
}
