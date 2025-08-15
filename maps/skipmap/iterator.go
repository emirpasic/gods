// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skipmap

import "github.com/emirpasic/gods/v2/containers"

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey[int, int] = (*Iterator[int, int])(nil)

// Iterator holding the iterator's state
type Iterator[K comparable, V any] struct {
	m       *Map[K, V]
	index   int
	element *element[K, V]
}

// Iterator returns a stateful iterator whose keys and values can be fetched by an index.
func (m *Map[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{m: m, index: -1, element: m.sentinel}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Next() bool {
	if iterator.index < iterator.m.size {
		iterator.index++
		iterator.element = iterator.element.nexts[0]
	}
	return iterator.element != iterator.m.sentinel
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
		iterator.element = iterator.element.prev
	}
	return iterator.element != iterator.m.sentinel
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Value() V {
	return iterator.element.value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Key() K {
	return iterator.element.key
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[K, V]) Begin() {
	iterator.index = -1
	iterator.element = iterator.m.sentinel
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[K, V]) End() {
	iterator.index = iterator.m.size
	iterator.element = iterator.m.sentinel
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool {
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}

// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool {
	for iterator.Prev() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}
