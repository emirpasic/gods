// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/maps"
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Map implementation
var _ maps.Map[string, int] = (*Map[string, int])(nil)

// Map holds the elements in a red-black tree
type Map[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
}

// New instantiates a tree map with the built-in comparator for K
func New[K cmp.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{tree: rbt.New[K, V]()}
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWith[K, V](comparator)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Min() (key K, value V, ok bool) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value, true
	}
	return key, value, false
}

// Max returns the maximum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Max() (key K, value V, ok bool) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value, true
	}
	return key, value, false
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
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, ok bool) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value, true
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
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, ok bool) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value, true
	}
	return foundKey, foundValue, false
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}
