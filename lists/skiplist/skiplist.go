// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skiplist

import (
	"fmt"
	"math/bits"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/emirpasic/gods/v2/lists"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert List implementation
var _ lists.List[int] = (*List[int])(nil)

const defaultSkipListHeight = 64

type (
	// List holds a sentinel element and the size of the list.
	List[T comparable] struct {
		sentinel *element[T]
		size     int
	}

	element[T comparable] struct {
		value T
		nexts []next[T]
		prev  *element[T]
	}

	next[T comparable] struct {
		next *element[T]
		dist int
	}
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[T comparable](ts ...T) *List[T] {
	var t T
	elem := &element[T]{
		value: t,
		nexts: make([]next[T], defaultSkipListHeight),
	}
	for i := range elem.nexts {
		elem.nexts[i] = next[T]{
			next: elem,
			dist: 1,
		}
	}
	elem.prev = elem
	s := &List[T]{sentinel: elem}
	s.Insert(0, ts...)
	return s
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (sl *List[T]) Add(values ...T) {
	sl.Insert(sl.size, values...)
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (sl *List[T]) Append(values ...T) {
	sl.Insert(sl.size, values...)
}

// Prepend prepends a values (or more)
func (sl *List[T]) Prepend(values ...T) {
	sl.Insert(0, values...)
}

func (sl *List[T]) getAt(index int) *element[T] {
	elem := sl.sentinel
	for height, i := len(elem.nexts)-1, 0; height >= 0; height-- {
		for i+elem.nexts[height].dist <= index+1 {
			i += elem.nexts[height].dist
			elem = elem.nexts[height].next
		}
	}
	return elem
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (sl *List[T]) Get(index int) (T, bool) {
	if index < 0 || index >= sl.size {
		var t T
		return t, false
	} else if index == 0 {
		return sl.sentinel.nexts[0].next.value, true
	} else if index == sl.size-1 {
		return sl.sentinel.prev.value, true
	}

	return sl.getAt(index).value, true
}

// Remove removes the element at the given index from the list.
func (sl *List[T]) Remove(index int) {
	if index < 0 || index >= sl.size {
		return
	}

	for height, i, prev := len(sl.sentinel.nexts)-1, 0, sl.sentinel; height >= 0; height-- {
		for i+prev.nexts[height].dist <= index {
			i += prev.nexts[height].dist
			prev = prev.nexts[height].next
		}

		if prev.nexts[height].dist > 1 {
			prev.nexts[height].dist--
		} else {
			prev.nexts[height].next.nexts[height].next.prev = prev
			prev.nexts[height] = next[T]{
				next: prev.nexts[height].next.nexts[height].next,
				dist: prev.nexts[height].dist + prev.nexts[height].next.nexts[height].dist - 1,
			}
		}
	}

	sl.size--
}

// Contains check if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (sl *List[T]) Contains(value ...T) bool {
	for _, v := range value {
		if _, ok := sl.IndexOf(v); !ok {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (sl *List[T]) Values() []T {
	values := make([]T, sl.size)
	elem := sl.sentinel
	for i := 0; i < sl.size; i++ {
		elem = elem.nexts[0].next
		values[i] = elem.value
	}
	return values
}

// IndexOf returns index of provided element
func (sl *List[T]) IndexOf(value T) (int, bool) {
	elem := sl.sentinel
	for i := 0; i < sl.size; i++ {
		elem = elem.nexts[0].next
		if elem.value == value {
			return i, true
		}
	}

	return -1, false
}

// Empty returns true if list does not contain any elements.
func (sl *List[T]) Empty() bool {
	return sl.size == 0
}

// Size returns number of elements within the list.
func (sl *List[T]) Size() int {
	return sl.size
}

// Clear removes all elements from the list.
func (sl *List[T]) Clear() {
	sl.size = 0
	for i := range sl.sentinel.nexts {
		sl.sentinel.nexts[i] = next[T]{
			next: sl.sentinel,
			dist: 1,
		}
	}
	sl.sentinel.prev = sl.sentinel
}

// Sort sorts values (in-place) using a Comparator.
func (sl *List[T]) Sort(comparator utils.Comparator[T]) {
	if sl.size < 2 {
		return
	}

	values := sl.Values()
	slices.SortFunc(values, comparator)

	elem := sl.sentinel
	for i := 0; i < sl.size; i++ {
		elem = elem.nexts[0].next
		elem.value = values[i]
	}
}

// Swap swaps values of two elements at the given indices.
func (sl *List[T]) Swap(a, b int) {
	elemA := sl.getAt(a)
	elemB := sl.getAt(b)
	elemA.value, elemB.value = elemB.value, elemA.value
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (sl *List[T]) Insert(index int, values ...T) {
	if index < 0 || index > sl.size {
		return
	}

	for _, value := range values {
		elem := &element[T]{
			value: value,
			nexts: make([]next[T], bits.TrailingZeros64(uint64(rand.Int64())<<1)),
		}

		prev := sl.sentinel
		for height, i := len(prev.nexts)-1, 0; height >= 0; height-- {
			for i+prev.nexts[height].dist <= index {
				i += prev.nexts[height].dist
				prev = prev.nexts[height].next
			}
			if height < len(elem.nexts) {
				elem.nexts[height] = next[T]{
					next: prev.nexts[height].next,
					dist: index - i + prev.nexts[height].dist,
				}
				prev.nexts[height] = next[T]{
					next: elem,
					dist: index - i + 1,
				}
			} else {
				prev.nexts[height].dist++
			}
		}
		elem.prev = prev
		elem.nexts[0].next.prev = elem
		sl.size++
		index++
	}
}

// Set value at specified index position
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (sl *List[T]) Set(index int, value T) {
	if index == sl.size {
		sl.Add(value)
		return
	}
	if index < 0 || index >= sl.size {
		return
	}
	if index == 0 {
		sl.sentinel.nexts[0].next.value = value
		return
	}
	if index == sl.size-1 {
		sl.sentinel.prev.value = value
		return
	}
	sl.getAt(index).value = value
}

// String returns a string representation of container
func (sl *List[T]) String() string {
	str := "SkipList\n"
	values := make([]string, sl.size)

	elem := sl.sentinel
	for i := 0; i < sl.size; i++ {
		elem = elem.nexts[0].next
		values[i] = fmt.Sprintf("%v", elem.value)
	}

	str += strings.Join(values, ", ")
	return str
}
