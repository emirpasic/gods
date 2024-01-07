// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"github.com/emirpasic/gods/v2/containers"
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
)

// Assert Enumerable implementation
var _ containers.EnumerableWithIndex[int] = (*Set[int])(nil)

// Each calls the given function once for each element, passing that element's index and value.
func (set *Set[T]) Each(f func(index int, value T)) {
	iterator := set.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (set *Set[T]) Map(f func(index int, value T) T) *Set[T] {
	newSet := &Set[T]{tree: rbt.NewWith[T, struct{}](set.tree.Comparator)}
	iterator := set.Iterator()
	for iterator.Next() {
		newSet.Add(f(iterator.Index(), iterator.Value()))
	}
	return newSet
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (set *Set[T]) Select(f func(index int, value T) bool) *Set[T] {
	newSet := &Set[T]{tree: rbt.NewWith[T, struct{}](set.tree.Comparator)}
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newSet.Add(iterator.Value())
		}
	}
	return newSet
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (set *Set[T]) Any(f func(index int, value T) bool) bool {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (set *Set[T]) All(f func(index int, value T) bool) bool {
	iterator := set.Iterator()
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
func (set *Set[T]) Find(f func(index int, value T) bool) (int, T) {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	var t T
	return -1, t
}
