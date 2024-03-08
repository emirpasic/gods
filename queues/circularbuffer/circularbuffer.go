// Copyright (c) 2021, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package circularbuffer implements the circular buffer.
//
// In computer science, a circular buffer, circular queue, cyclic buffer or ring buffer is a data structure that uses a single, fixed-size buffer as if it were connected end-to-end. This structure lends itself easily to buffering data streams.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Circular_buffer
package circularbuffer

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/queues"
)

// Assert Queue implementation
var _ queues.Queue[int] = (*Queue[int])(nil)

// Queue holds values in a slice.
type Queue[T any] struct {
	values  []T
	start   int
	end     int
	full    bool
	maxSize int
	size    int
}

// New instantiates a new empty queue with the specified size of maximum number of elements that it can hold.
// This max size of the buffer cannot be changed.
func New[T any](maxSize int) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}
	queue := &Queue[T]{maxSize: maxSize}
	queue.Clear()
	return queue
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[T]) Enqueue(value T) {
	if queue.Full() {
		queue.Dequeue()
	}
	queue.values[queue.end] = value
	queue.end = queue.end + 1
	if queue.end >= queue.maxSize {
		queue.end = 0
	}
	if queue.end == queue.start {
		queue.full = true
	}

	queue.size = queue.calculateSize()
}

// Dequeue removes first element of the queue and returns it, or the 0-value if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	if queue.Empty() {
		return value, false
	}

	value, ok = queue.values[queue.start], true
	queue.start = queue.start + 1
	if queue.start >= queue.maxSize {
		queue.start = 0
	}
	queue.full = false
	queue.size = queue.size - 1

	return
}

// Peek returns first element of the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	if queue.Empty() {
		return value, false
	}
	return queue.values[queue.start], true
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[T]) Empty() bool {
	return queue.Size() == 0
}

// Full returns true if the queue is full, i.e. has reached the maximum number of elements that it can hold.
func (queue *Queue[T]) Full() bool {
	return queue.Size() == queue.maxSize
}

// Size returns number of elements within the queue.
func (queue *Queue[T]) Size() int {
	return queue.size
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.values = make([]T, queue.maxSize, queue.maxSize)
	queue.start = 0
	queue.end = 0
	queue.full = false
	queue.size = 0
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[T]) Values() []T {
	values := make([]T, queue.Size(), queue.Size())
	for i := 0; i < queue.Size(); i++ {
		values[i] = queue.values[(queue.start+i)%queue.maxSize]
	}
	return values
}

// String returns a string representation of container
func (queue *Queue[T]) String() string {
	str := "CircularBuffer\n"
	var values []string
	for _, value := range queue.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (queue *Queue[T]) withinRange(index int) bool {
	return index >= 0 && index < queue.size
}

func (queue *Queue[T]) calculateSize() int {
	if queue.end < queue.start {
		return queue.maxSize - queue.start + queue.end
	} else if queue.end == queue.start {
		if queue.full {
			return queue.maxSize
		}
		return 0
	}
	return queue.end - queue.start
}
