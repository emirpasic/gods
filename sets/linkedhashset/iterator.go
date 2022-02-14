// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"container/list"

	"github.com/emirpasic/gods/containers"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	next    *list.Element
	current *list.Element
	prev    *list.Element
	src     *list.List
	index   int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (set *Set) Iterator() Iterator {
	return Iterator{
		next:  set.ordering.Front(),
		src:   &set.ordering,
		index: -1,
	}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	if iterator.next == nil {
		return false
	}
	iterator.prev = iterator.current
	iterator.current = iterator.next
	iterator.next = iterator.next.Next()
	iterator.index++
	return true
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Prev() bool {
	if iterator.prev == nil {
		return false
	}
	iterator.next = iterator.current
	iterator.current = iterator.prev
	iterator.prev = iterator.prev.Prev()
	iterator.index--
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	if iterator.current == nil {
		return nil
	}
	return iterator.current.Value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Begin() {
	iterator.prev = nil
	iterator.next = iterator.src.Front()
	iterator.current = nil
	iterator.index = -1
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator) End() {
	iterator.next = nil
	iterator.prev = iterator.src.Back()
	iterator.current = nil
	iterator.index = iterator.src.Len()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) First() bool {
	iterator.prev = nil
	iterator.current = iterator.src.Front()
	if iterator.current != nil {
		iterator.next = iterator.current.Next()
	} else {
		iterator.next = nil
	}
	iterator.index = 0
	return iterator.current != nil
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Last() bool {
	iterator.next = nil
	iterator.current = iterator.src.Back()
	if iterator.current != nil {
		iterator.prev = iterator.current.Prev()
	} else {
		iterator.prev = nil
	}
	iterator.index = iterator.src.Len() - 1
	return iterator.current != nil
}
