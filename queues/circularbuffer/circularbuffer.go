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
type Queue[T comparable] struct {
	values []T // maintain this slices's capacity equal to its length
	start  int
	end    int
	size   int
}

// New instantiates a new empty queue with the specified size of maximum number of elements that it can hold.
// This max size of the buffer cannot be changed.
func New[T comparable](maxSize int) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}
	return &Queue[T]{
		values: make([]T, maxSize),
	}
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[T]) Enqueue(value T) {
	if queue.Full() {
		queue.start = queue.start + 1
		if queue.start == len(queue.values) {
			queue.start = 0
		}
		queue.size--
	}
	queue.values[queue.end] = value
	queue.end = queue.end + 1
	if queue.end >= len(queue.values) {
		queue.end = 0
	}
	queue.size++
}

// Dequeue removes first element of the queue and returns it, or the 0-value if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	if queue.Empty() {
		return value, false
	}

	value, ok = queue.values[queue.start], true
	queue.start = queue.start + 1
	if queue.start >= len(queue.values) {
		queue.start = 0
	}
	queue.size--

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
	return queue.size == 0
}

// Full returns true if the queue is full, i.e. has reached the maximum number of elements that it can hold.
func (queue *Queue[T]) Full() bool {
	return queue.size == len(queue.values)
}

// Size returns number of elements within the queue.
func (queue *Queue[T]) Size() int {
	return queue.size
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	clear(queue.values)
	queue.start = 0
	queue.end = 0
	queue.size = 0
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[T]) Values() []T {
	values := make([]T, queue.size)
	for i := 0; i < queue.size; i++ {
		values[i] = queue.values[(queue.start+i)%len(queue.values)]
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
