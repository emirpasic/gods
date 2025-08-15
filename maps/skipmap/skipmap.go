// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skipmap

import (
	"cmp"
	"fmt"
	"math/bits"
	"math/rand/v2"
	"strings"

	"github.com/emirpasic/gods/v2/maps"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Map implementation
var _ maps.Map[int, int] = (*Map[int, int])(nil)

const defaultSkipMapHeight = 64

type (
	// Map holds a sentinel element and the size of the map.
	// Keys are kept in sorted order and duplicate keys are not allowed.
	Map[K comparable, V any] struct {
		sentinel   *element[K, V]
		size       int
		comparator utils.Comparator[K]
	}

	element[K comparable, V any] struct {
		key   K
		value V
		nexts []*element[K, V]
		prev  *element[K, V]
	}
)

// New instantiates a new map with the default comparator
func New[K cmp.Ordered, V any]() *Map[K, V] {
	return NewWith[K, V](cmp.Compare[K])
}

// New instantiates a new map with the provided comparator
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Map[K, V] {
	var k K
	var v V
	elem := &element[K, V]{
		key:   k,
		value: v,
		nexts: make([]*element[K, V], defaultSkipMapHeight),
	}
	for i := range elem.nexts {
		elem.nexts[i] = elem
	}
	m := &Map[K, V]{
		sentinel:   elem,
		comparator: comparator,
	}
	elem.prev = elem
	return m
}

func (m *Map[K, V]) getAt(key K) *element[K, V] {
	elem := m.sentinel
	for height := len(elem.nexts) - 1; height >= 0; height-- {
		for elem.nexts[height] != m.sentinel && m.comparator(elem.nexts[height].key, key) <= 0 {
			elem = elem.nexts[height]
		}
	}

	if elem != m.sentinel && m.comparator(elem.key, key) == 0 {
		return elem
	}
	return nil
}

// Put inserts key-value pair into the map. If the key already exists, the value is updated.
func (m *Map[K, V]) Put(key K, value V) {
	if elem := m.getAt(key); elem != nil {
		elem.value = value
		return
	}

	elem := &element[K, V]{
		key:   key,
		value: value,
		nexts: make([]*element[K, V], bits.TrailingZeros64(uint64(rand.Int64())<<1)),
	}

	prev := m.sentinel
	for level := len(prev.nexts) - 1; level >= 0; level-- {
		for prev.nexts[level] != m.sentinel && m.comparator(prev.nexts[level].key, key) < 0 {
			prev = prev.nexts[level]
		}

		if level < len(elem.nexts) {
			elem.nexts[level] = prev.nexts[level]
			prev.nexts[level] = elem
		}
	}
	elem.prev = prev
	elem.nexts[0].prev = elem
	m.size++
}

// Get returns the value for the given key and whether the key exists in the map
func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	if elem := m.getAt(key); elem != nil {
		return elem.value, true
	}
	return value, false
}

// Remove removes the key-value pair for the given key from the map
func (m *Map[K, V]) Remove(key K) {
	if !m.ContainsKey(key) {
		return
	}

	for height, prev := len(m.sentinel.nexts)-1, m.sentinel; height >= 0; height-- {
		for prev.nexts[height] != m.sentinel && m.comparator(prev.nexts[height].key, key) < 0 {
			prev = prev.nexts[height]
		}

		// If this level points to the element we want to remove, update the pointer
		if prev.nexts[height] != m.sentinel && m.comparator(prev.nexts[height].key, key) == 0 {
			toRemove := prev.nexts[height]
			prev.nexts[height] = toRemove.nexts[height]

			// Update prev pointer on level 0
			if height == 0 {
				if toRemove.nexts[0] != m.sentinel {
					toRemove.nexts[0].prev = prev
				} else {
					m.sentinel.prev = prev
				}
			}
		}
	}
	m.size--
}

// ContainsKey checks if the given key exists in the map
func (m *Map[K, V]) ContainsKey(key K) bool {
	_, found := m.Get(key)
	return found
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.size == 0
}

// Size returns number of elements within the map
func (m *Map[K, V]) Size() int {
	return m.size
}

// Clear removes all elements from the map
func (m *Map[K, V]) Clear() {
	m.size = 0
	for i := range m.sentinel.nexts {
		m.sentinel.nexts[i] = m.sentinel
	}
	m.sentinel.prev = m.sentinel
}

// Min returns the minimum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Min() (key K, value V, ok bool) {
	if m.Empty() {
		return key, value, false
	}

	return m.sentinel.nexts[0].key, m.sentinel.nexts[0].value, true
}

// Max returns the maximum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Max() (key K, value V, ok bool) {
	if m.Empty() {
		return key, value, false
	}

	return m.sentinel.prev.key, m.sentinel.prev.value, true
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, found bool) {
	current := m.sentinel
	for level := len(m.sentinel.nexts) - 1; level >= 0; level-- {
		for current.nexts[level] != m.sentinel && m.comparator(current.nexts[level].key, key) <= 0 {
			current = current.nexts[level]
		}
	}

	if current != m.sentinel {
		return current.key, current.value, true
	}
	return foundKey, foundValue, false
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, found bool) {
	current := m.sentinel
	for level := len(m.sentinel.nexts) - 1; level >= 0; level-- {
		for current.nexts[level] != m.sentinel && m.comparator(current.nexts[level].key, key) < 0 {
			current = current.nexts[level]
		}
	}
	current = current.nexts[0]
	if current != m.sentinel {
		return current.key, current.value, true
	}
	return foundKey, foundValue, false
}

// Keys returns all keys in the map as a slice in sorted order
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.size)
	elem := m.sentinel
	for i := 0; i < m.size; i++ {
		elem = elem.nexts[0]
		keys[i] = elem.key
	}
	return keys
}

// Values returns all values in the map as a slice in key-sorted order
func (m *Map[K, V]) Values() []V {
	values := make([]V, m.size)
	elem := m.sentinel
	for i := 0; i < m.size; i++ {
		elem = elem.nexts[0]
		values[i] = elem.value
	}
	return values
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "SkipMap\nmap["

	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
