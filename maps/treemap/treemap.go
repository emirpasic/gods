/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

// Implementation of order map backed by red-black tree.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Associative_array

package treemap

import (
	"github.com/emirpasic/gods/maps"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

func assertInterfaceImplementation() {
	var _ maps.Interface = (*Map)(nil)
}

type Map struct {
	tree *rbt.Tree
}

// Instantiates a tree map with the custom comparator.
func NewWith(comparator utils.Comparator) *Map {
	return &Map{tree: rbt.NewWith(comparator)}
}

// Instantiates a tree map with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Map {
	return &Map{tree: rbt.NewWithIntComparator()}
}

// Instantiates a tree map with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Map {
	return &Map{tree: rbt.NewWithStringComparator()}
}

// Inserts node into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Put(key interface{}, value interface{}) {
	m.tree.Put(key, value)
}

// Searches the node in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	return m.tree.Get(key)
}

// Remove the node from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Remove(key interface{}) {
	m.tree.Remove(key)
}

// Returns true if map does not contain any nodes
func (m *Map) Empty() bool {
	return m.tree.Empty()
}

// Returns number of items in the map.
func (m *Map) Size() int {
	return m.tree.Size()
}

// Returns all keys in-order
func (m *Map) Keys() []interface{} {
	return m.tree.Keys()
}

// Returns all values in-order based on the key.
func (m *Map) Values() []interface{} {
	return m.tree.Values()
}

// Removes all nodes from the map.
func (m *Map) Clear() {
	m.tree.Clear()
}

func (m *Map) String() string {
	str := "TreeMap\n"
	str += m.tree.String()
	return str
}
