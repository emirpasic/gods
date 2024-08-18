// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashmap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/maps"
)

// Assert Map implementation
var _ maps.Map[string, int] = (*Map[string, int])(nil)

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[K comparable, V any] struct {
	table map[K]*element[K, V]
	first *element[K, V]
	last  *element[K, V]
}

type element[K comparable, V any] struct {
	key   K
	value V
	prev  *element[K, V]
	next  *element[K, V]
}

// New instantiates a linked-hash-map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		table: make(map[K]*element[K, V]),
	}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	if el, contains := m.table[key]; contains {
		el.value = value
	} else {
		e := &element[K, V]{
			key:   key,
			value: value,
			prev:  m.last,
		}

		if m.Size() == 0 {
			m.first = e
			m.last = e
		} else {
			m.last.next = e
			m.last = e
		}
		m.table[key] = e
	}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	element := m.table[key]
	if element != nil {
		found = true
		value = element.value
	}
	return
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	if element, contains := m.table[key]; contains {
		if element == m.first {
			m.first = element.next
		}
		if element == m.last {
			m.last = element.prev
		}
		if element.prev != nil {
			element.prev.next = element.next
		}
		if element.next != nil {
			element.next.prev = element.prev
		}
		element = nil
		delete(m.table, key)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return len(m.table)
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		keys[count] = it.Key()
		count++
	}
	return keys
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.table = make(map[K]*element[K, V])
	m.first = nil
	m.last = nil
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "LinkedHashMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
