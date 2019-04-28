package xtypes

import (
	"testing"
)

func TestNewSemaphore(t *testing.T) {
	size := 1
	sema := NewSemaphore(size)
	if sema == nil {
		t.Fatal("failed to create semaphore")
	}

	if actual := cap(sema); actual != size {
		t.Fatalf("expected semaphore size is %d, actual is %d", size, actual)
	}
}

func TestSemaphore_Acquire(t *testing.T) {
	sema := NewSemaphore(2)

	sema.Acquire(1)

	expected := 1
	if v := len(sema); v != expected {
		t.Fatalf("expected %d type, got %d", expected, v)
	}
}

func TestSemaphore_Release(t *testing.T) {
	sema := NewSemaphore(2)

	sema.Acquire(1)
	sema.Release(1)

	expected := 0
	if v := len(sema); v != expected {
		t.Fatalf("expected %d type, got %d", expected, v)
	}
}

func TestSemaphore_ReleaseEmpty(t *testing.T) {
	sema := NewSemaphore(1)

	sema.Acquire(1)
	sema.Release(1)

	sema.Release(1)

	expected := 0
	if v := len(sema); v != expected {
		t.Fatalf("expected %d type, got %d", expected, v)
	}
}
