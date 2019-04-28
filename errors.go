package xtypes

import (
	"errors"
)

var (
	// ErrEmptyQueue is returned when an non-applicable operation was called on queue due to its empty state.
	ErrEmptyQueue = errors.New("empty queue")

	// ErrInvalidQueue is returned when an non-applicable operation was called on a nil queue.
	ErrInvalidQueue = errors.New("invalid queue")
)
