// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import "github.com/emirpasic/gods/v2/containers"

// Assert Enumerable implementation
var _ containers.EnumerableWithIndex[int] = (*List[int])(nil)

// Each calls the given function once for each element, passing that element's index and value.
func (list *List[T]) Each(f func(index int, value T)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (list *List[T]) Map(f func(index int, value T) T) *List[T] {
	newList := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (list *List[T]) Select(f func(index int, value T) bool) *List[T] {
	newList := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (list *List[T]) Any(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (list *List[T]) All(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (list *List[T]) Find(f func(index int, value T) bool) (index int, value T) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, value
}
