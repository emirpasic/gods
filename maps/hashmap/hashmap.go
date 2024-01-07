// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"

	"github.com/emirpasic/gods/v2/maps"
)

// Assert Map implementation
var _ maps.Map[string, int] = (*Map[string, int])(nil)

// Map holds the elements in go's native map
type Map[K comparable, V any] struct {
	m map[K]V
}

// New instantiates a hash map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	clear(m.m)
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
