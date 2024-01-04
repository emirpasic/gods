// Copyright (c) 2015, Emir Pasic & Eren Dursun. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"github.com/emirpasic/gods/containers"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)

// Iterator holding the iterator's state
type Iterator struct {
	m       *Map
	index   int
	element *element
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (m *Map) Iterator() Iterator {
	return Iterator{
		m:     m,
		index: -1,
	}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.m.Size() {
		iterator.index++
	}
	if !iterator.m.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index != 0 {
		iterator.element = iterator.element.next
	} else {
		iterator.element = iterator.m.first
	}
	return true
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	if !iterator.m.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == iterator.m.Size()-1 {
		iterator.element = iterator.m.last
	} else {
		iterator.element = iterator.element.prev
	}
	return iterator.m.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	if iterator.element != nil {
		return iterator.element.value
	}
	return nil
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator) Key() interface{} {
	return iterator.element.key
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator) End() {
	iterator.index = iterator.m.Size()
	iterator.element = nil
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) NextTo(f func(key interface{}, value interface{}) bool) bool {
	for iterator.Next() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}

// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) PrevTo(f func(key interface{}, value interface{}) bool) bool {
	for iterator.Prev() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}

// Check that the index is within bounds of the list
func (m *Map) withinRange(index int) bool {
	return index >= 0 && index < m.Size()
}
