// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skipmap

import "github.com/emirpasic/gods/v2/containers"

// Assert Enumerable implementation
var _ containers.EnumerableWithKey[int, int] = (*Map[int, int])(nil)

// Each calls the given function once for each element, passing that element's key and value.
func (m *Map[K, V]) Each(f func(key K, value V)) {
	iterator := m.Iterator()
	for iterator.Next() {
		f(iterator.Key(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (m *Map[K, V]) Map(f func(key K, value V) (K, V)) *Map[K, V] {
	newMap := NewWith[K, V](m.comparator)
	iterator := m.Iterator()
	for iterator.Next() {
		key, value := f(iterator.Key(), iterator.Value())
		newMap.Put(key, value)
	}
	return newMap
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V] {
	newMap := NewWith[K, V](m.comparator)
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			newMap.Put(iterator.Key(), iterator.Value())
		}
	}
	return newMap
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (m *Map[K, V]) All(f func(key K, value V) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if !f(iterator.Key(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil otherwise
// if no element matches the criteria.
func (m *Map[K, V]) Find(f func(key K, value V) bool) (key K, value V) {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return iterator.Key(), iterator.Value()
		}
	}
	return key, value
}
