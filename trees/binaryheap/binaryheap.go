/*
Copyright (c) 2015, Emir Pasic
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Implementation of binary heap backed by ArrayList.
// Comparator defines this heap as either min or max heap.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Binary_heap

package binaryheap

import (
	"fmt"
	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
	"strings"
)

func assertInterfaceImplementation() {
	var _ trees.Tree = (*Heap)(nil)
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

type Heap struct {
	list       *arraylist.List
	Comparator utils.Comparator
}

// Instantiates a new empty heap tree with the custom comparator.
func NewWith(comparator utils.Comparator) *Heap {
	return &Heap{list: arraylist.New(), Comparator: comparator}
}

// Instantiates a new empty heap with the IntComparator, i.e. elements are of type int.
func NewWithIntComparator() *Heap {
	return &Heap{list: arraylist.New(), Comparator: utils.IntComparator}
}

// Instantiates a new empty heap with the StringComparator, i.e. elements are of type string.
func NewWithStringComparator() *Heap {
	return &Heap{list: arraylist.New(), Comparator: utils.StringComparator}
}

// Pushes a value onto the heap and bubbles it up accordingly.
func (heap *Heap) Push(value interface{}) {
	heap.list.Add(value)
	heap.bubbleUp()
}

// Pops (removes) top element on heap and returns it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to pop.
func (heap *Heap) Pop() (value interface{}, ok bool) {
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

// Returns top element on the heap without removing it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to peek.
func (heap *Heap) Peek() (value interface{}, ok bool) {
	return heap.list.Get(0)
}

// Returns true if heap does not contain any elements.
func (heap *Heap) Empty() bool {
	return heap.list.Empty()
}

// Returns number of elements within the heap.
func (heap *Heap) Size() int {
	return heap.list.Size()
}

// Removes all elements from the heap.
func (heap *Heap) Clear() {
	heap.list.Clear()
}

// Returns all elements in the heap.
func (heap *Heap) Values() []interface{} {
	return heap.list.Values()
}

type Iterator struct {
	heap  *Heap
	index int
}

// Returns a stateful iterator whose values can be fetched by an index.
func (heap *Heap) Iterator() Iterator {
	return Iterator{heap: heap, index: -1}
}

// Moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	iterator.index += 1
	return iterator.heap.withinRange(iterator.index)
}

// Returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.heap.list.Get(iterator.index)
	return value
}

// Returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

func (heap *Heap) String() string {
	str := "BinaryHeap\n"
	values := []string{}
	for _, value := range heap.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Performs the "bubble down" operation. This is to place the element that is at the
// root of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap) bubbleDown() {
	index := 0
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
func (heap *Heap) bubbleUp() {
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

// Check that the index is withing bounds of the list
func (heap *Heap) withinRange(index int) bool {
	return index >= 0 && index < heap.list.Size()
}
