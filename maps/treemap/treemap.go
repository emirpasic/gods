/*
Copyright (c) 2015, Emir Pasic
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Implementation of order map backed by red-black tree.
// Elements are ordered by key in the map.
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

// Inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Put(key interface{}, value interface{}) {
	m.tree.Put(key, value)
}

// Searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	return m.tree.Get(key)
}

// Remove the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Remove(key interface{}) {
	m.tree.Remove(key)
}

// Returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.tree.Empty()
}

// Returns number of elements in the map.
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

// Removes all elements from the map.
func (m *Map) Clear() {
	m.tree.Clear()
}

// Returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map) Min() (key interface{}, value interface{}) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	return nil, nil
}

// Returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map) Max() (key interface{}, value interface{}) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	return nil, nil
}

func (m *Map) String() string {
	str := "TreeMap\n"
	str += m.tree.String()
	return str
}
