// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binaryheap implements a binary heap backed by array list.
//
// Comparator defines this heap as either min or max heap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Binary_heap
package binaryheap

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Tree implementation
var _ trees.Tree[int] = (*Heap[int])(nil)

// Heap holds elements in an array-list
type Heap[T comparable] struct {
	list       *arraylist.List[T]
	Comparator utils.Comparator[T]
}

// New instantiates a new empty heap tree with the built-in comparator for T
func New[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{list: arraylist.New[T](), Comparator: cmp.Compare[T]}
}

// NewWith instantiates a new empty heap tree with the custom comparator.
func NewWith[T comparable](comparator utils.Comparator[T]) *Heap[T] {
	return &Heap[T]{list: arraylist.New[T](), Comparator: comparator}
}

// Push adds a value onto the heap and bubbles it up accordingly.
func (heap *Heap[T]) Push(values ...T) {
	if len(values) == 1 {
		heap.list.Add(values[0])
		heap.bubbleUp()
	} else {
		// Reference: https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap
		for _, value := range values {
			heap.list.Add(value)
		}
		size := heap.list.Size()/2 + 1
		for i := size; i >= 0; i-- {
			heap.bubbleDownIndex(i)
		}
	}
}

// Pop removes top element on heap and returns it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to pop.
func (heap *Heap[T]) Pop() (value T, ok bool) {
	value, ok = heap.list.Get(0)
	if !ok {
		return
	}
	lastIndex := heap.list.Size() - 1
	heap.list.Swap(0, lastIndex)
	heap.list.Remove(lastIndex)
	heap.bubbleDown()
	return
}

// Peek returns top element on the heap without removing it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to peek.
func (heap *Heap[T]) Peek() (value T, ok bool) {
	return heap.list.Get(0)
}

// Empty returns true if heap does not contain any elements.
func (heap *Heap[T]) Empty() bool {
	return heap.list.Empty()
}

// Size returns number of elements within the heap.
func (heap *Heap[T]) Size() int {
	return heap.list.Size()
}

// Clear removes all elements from the heap.
func (heap *Heap[T]) Clear() {
	heap.list.Clear()
}

// Values returns all elements in the heap.
func (heap *Heap[T]) Values() []T {
	values := make([]T, heap.list.Size(), heap.list.Size())
	for it := heap.Iterator(); it.Next(); {
		values[it.Index()] = it.Value()
	}
	return values
}

// String returns a string representation of container
func (heap *Heap[T]) String() string {
	str := "BinaryHeap\n"
	values := []string{}
	for it := heap.Iterator(); it.Next(); {
		values = append(values, fmt.Sprintf("%v", it.Value()))
	}
	str += strings.Join(values, ", ")
	return str
}

// Performs the "bubble down" operation. This is to place the element that is at the root
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleDown() {
	heap.bubbleDownIndex(0)
}

// Performs the "bubble down" operation. This is to place the element that is at the index
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleDownIndex(index int) {
	size := heap.list.Size()
	for leftIndex := index<<1 + 1; leftIndex < size; leftIndex = index<<1 + 1 {
		rightIndex := index<<1 + 2
		smallerIndex := leftIndex
		leftValue, _ := heap.list.Get(leftIndex)
		rightValue, _ := heap.list.Get(rightIndex)
		if rightIndex < size && heap.Comparator(leftValue, rightValue) > 0 {
			smallerIndex = rightIndex
		}
		indexValue, _ := heap.list.Get(index)
		smallerValue, _ := heap.list.Get(smallerIndex)
		if heap.Comparator(indexValue, smallerValue) > 0 {
			heap.list.Swap(index, smallerIndex)
		} else {
			break
		}
		index = smallerIndex
	}
}

// Performs the "bubble up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place so that
// the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleUp() {
	index := heap.list.Size() - 1
	for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 {
		indexValue, _ := heap.list.Get(index)
		parentValue, _ := heap.list.Get(parentIndex)
		if heap.Comparator(parentValue, indexValue) <= 0 {
			break
		}
		heap.list.Swap(index, parentIndex)
		index = parentIndex
	}
}

// Check that the index is within bounds of the list
func (heap *Heap[T]) withinRange(index int) bool {
	return index >= 0 && index < heap.list.Size()
}
